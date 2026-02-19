package discovery

import (
	"sync"
	"time"
)

// -------- Node Discovery & Health --------
// In a distributed system, nodes need to find each other and detect failures.

// NodeStatus represents whether a node is alive or suspected dead.
type NodeStatus int

const (
	StatusAlive NodeStatus = iota
	StatusSuspect
	StatusDead
)

// Node holds metadata about a single cache node in the cluster.
//
//	we received a heartbeat from it.
type Node struct {
  Addr string
  CurrStatus NodeStatus
  LastHB time.Time
}

// Registry keeps track of all known nodes and their health.
//
//	a timeout after which a node is considered dead.
type Registry struct {
  AddrNode map[string]Node
  mu sync.Mutex
  timeOut time.Duration
}

// NewRegistry creates a Registry that marks nodes dead after the given timeout.
func NewRegistry(healthTimeout time.Duration) *Registry {
  r := &Registry{
    AddrNode: make(map[string]Node),
    timeOut: healthTimeout,
  }

  ticker := time.NewTicker(healthTimeout)
  go func() {
    for {
      select {
        case <- ticker.C:
          r.checkHealth()
      }
    }
  }()
  return r
}

// Register adds a new node to the cluster or updates an existing one's heartbeat.
func (r *Registry) Register(addr string) {
  r.mu.Lock()
  defer r.mu.Unlock()

  node, ok := r.AddrNode[addr]
  if !ok {
    r.AddrNode[addr] = Node{
      Addr: addr,
      CurrStatus: StatusAlive,
      LastHB: time.Now(),
    }
  } else {
    node.LastHB = time.Now()
    r.AddrNode[addr] = node
  }
}

// Heartbeat updates the last-seen time for a node.
func (r *Registry) Heartbeat(addr string) {
  r.mu.Lock()
  defer r.mu.Unlock()

  node, ok := r.AddrNode[addr]
  if !ok {
    return
  }
  
  node.LastHB = time.Now()
  if node.CurrStatus == StatusSuspect {
    node.CurrStatus = StatusAlive
  }

  r.AddrNode[addr] = node
}

// Unregister removes a node from the cluster.
func (r *Registry) Unregister(addr string) {
  r.mu.Lock()
  defer r.mu.Unlock()

  delete(r.AddrNode, addr)
}

// AliveNodes returns the addresses of all nodes currently StatusAlive.
func (r *Registry) AliveNodes() []string {
  r.mu.Lock()
  defer r.mu.Unlock()

  addr := []string{}
  for add, node := range r.AddrNode {
    if node.CurrStatus == StatusAlive {
      addr = append(addr, add)
    }
  }
  return addr
}

// checkHealth iterates over all nodes and marks those with stale heartbeats
// as StatusSuspect or StatusDead.
//
//	If last heartbeat > 2*timeout, mark StatusDead and remove.
func (r *Registry) checkHealth() {
  r.mu.Lock()
  defer r.mu.Unlock()

  for add, node := range r.AddrNode {
    if time.Since(node.LastHB) > r.timeOut {
      node.CurrStatus = StatusSuspect
      r.AddrNode[add] = node
    } 
    if time.Since(node.LastHB) > 2*r.timeOut {
      node.CurrStatus = StatusDead
      delete(r.AddrNode, add)
    }
  }
}
