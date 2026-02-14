package discovery

import "time"

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
// TODO: Store the node's address, its current status, and the last time
//       we received a heartbeat from it.
type Node struct {
	// YOUR CODE HERE
}

// Registry keeps track of all known nodes and their health.
// TODO: Store a map of address→Node, protect with mutex, and configure
//       a timeout after which a node is considered dead.
type Registry struct {
	// YOUR CODE HERE
}

// NewRegistry creates a Registry that marks nodes dead after the given timeout.
// TODO: Start a background goroutine that periodically checks heartbeat timestamps.
func NewRegistry(healthTimeout time.Duration) *Registry {
	// YOUR CODE HERE
	return nil
}

// Register adds a new node to the cluster or updates an existing one's heartbeat.
func (r *Registry) Register(addr string) {
	// YOUR CODE HERE
}

// Heartbeat updates the last-seen time for a node.
// TODO: If the node was StatusSuspect, move it back to StatusAlive.
func (r *Registry) Heartbeat(addr string) {
	// YOUR CODE HERE
}

// Unregister removes a node from the cluster.
func (r *Registry) Unregister(addr string) {
	// YOUR CODE HERE
}

// AliveNodes returns the addresses of all nodes currently StatusAlive.
func (r *Registry) AliveNodes() []string {
	// YOUR CODE HERE
	return nil
}

// checkHealth iterates over all nodes and marks those with stale heartbeats
// as StatusSuspect or StatusDead.
// TODO: If last heartbeat > timeout, mark StatusSuspect.
//       If last heartbeat > 2*timeout, mark StatusDead and remove.
func (r *Registry) checkHealth() {
	// YOUR CODE HERE
}
