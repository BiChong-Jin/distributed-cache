package consistent

import (
	"fmt"
	"hash"
	"hash/crc32"
	"sort"
	"sync"
)

// -------- Consistent Hashing --------
// This is how you decide WHICH node owns a given key.
// Without this, adding/removing a node would remap almost all keys.
// With consistent hashing, only ~1/N keys get remapped.

// HashRing distributes keys across nodes using consistent hashing.
type HashRing struct {
	mu       sync.RWMutex
	//   - hashes []int           → sorted ring positions
	//   - ring   map[int]string  → hash position → node name
	//   - replicas int           → number of virtual nodes per real node
	//   - hasher hash.Hash32     → hash function (e.g. crc32, fnv)
  hashes []int
  replicas int
  ring map[int]string
  hasher hash.Hash32
}

// Ensure the compiler knows we use hash.Hash (suppress unused import).
var _ hash.Hash

// NewHashRing creates a ring with the given number of virtual nodes per real node.
func NewHashRing(replicas int) *HashRing {
	return &HashRing{
    replicas: replicas,
    ring: make(map[int]string),
    hasher: crc32.NewIEEE(),
  }
}

// AddNode adds a real node to the ring.
func (h *HashRing) AddNode(addr string) {
  h.mu.Lock()
  defer h.mu.Unlock()

  replicas := h.replicas
  for i := 0; i < replicas; i++ {
    replicaAddr := fmt.Sprintf("%s-%d", addr, i)
    h.hasher.Reset()
    h.hasher.Write([]byte(replicaAddr))
    hashReplicaValue := int(h.hasher.Sum32())
    h.ring[hashReplicaValue] = addr
    h.hashes = append(h.hashes, hashReplicaValue)
  }

  sort.Ints(h.hashes)
}

// RemoveNode removes a node and all its virtual nodes from the ring.
func (h *HashRing) RemoveNode(addr string) {
  h.mu.Lock()
  defer h.mu.Unlock()

  for key, val := range h.ring {
    if val == addr {
      delete(h.ring, key)
    }
  }

  h.hashes = []int{}
  for keys := range h.ring {
    h.hashes = append(h.hashes, keys)
  }

  sort.Ints(h.hashes)
}

// GetNode returns the node responsible for the given key.
func (h *HashRing) GetNode(key string) string {
  h.mu.RLock()
  defer h.mu.RUnlock()

  if len(h.hashes) == 0 {
    return  ""
  }

  h.hasher.Reset()
  h.hasher.Write([]byte(key))
  hashKey := int(h.hasher.Sum32())
  
  idx := sort.Search(len(h.hashes), func(i int) bool {
    return h.hashes[i] >= hashKey
  })
  
  if idx == len(h.hashes) {
    idx = 0
  }
  
  nn := h.ring[h.hashes[idx]]

	return nn
}

// GetNodes returns all unique real node addresses currently in the ring.
func (h *HashRing) GetNodes() []string {
  h.mu.RLock()
  defer h.mu.RUnlock()

  set := make(map[string]bool)
  for _, v := range h.ring {
    set[v] = true
  }

  addrs := []string{}
  for k, _ := range set {
    addrs = append(addrs, k)
  }
  
	return addrs
}
