package consistent

import (
	"hash"
	"sync"
)

// -------- Consistent Hashing --------
// This is how you decide WHICH node owns a given key.
// Without this, adding/removing a node would remap almost all keys.
// With consistent hashing, only ~1/N keys get remapped.

// HashRing distributes keys across nodes using consistent hashing.
// TODO: Store a sorted slice of hash values, a map from hash→node name,
//       the number of virtual nodes (replicas) per real node, and a hash function.
type HashRing struct {
	mu       sync.RWMutex
	// YOUR CODE HERE
	//
	// Hints:
	//   - hashes []int           → sorted ring positions
	//   - ring   map[int]string  → hash position → node name
	//   - replicas int           → number of virtual nodes per real node
	//   - hasher hash.Hash32     → hash function (e.g. crc32, fnv)
}

// Ensure the compiler knows we use hash.Hash (suppress unused import).
var _ hash.Hash

// NewHashRing creates a ring with the given number of virtual nodes per real node.
// TODO: Initialize the struct. A good default for replicas is 100-150.
func NewHashRing(replicas int) *HashRing {
	// YOUR CODE HERE
	return nil
}

// AddNode adds a real node to the ring.
// TODO: For each replica (0..replicas-1), hash a string like "nodeAddr-0", "nodeAddr-1", etc.
//       Store each hash in the sorted slice and map it to the node address.
//       Keep the hashes slice sorted (use sort.Ints).
func (h *HashRing) AddNode(addr string) {
	// YOUR CODE HERE
}

// RemoveNode removes a node and all its virtual nodes from the ring.
// TODO: Recompute the hashes for each replica of this node, remove them
//       from the slice and the map. Rebuild the sorted slice.
func (h *HashRing) RemoveNode(addr string) {
	// YOUR CODE HERE
}

// GetNode returns the node responsible for the given key.
// TODO: Hash the key, binary search (sort.Search) for the first ring position >= hash.
//       If past the end, wrap around to index 0 (that's the "ring" part).
//       Return the node name from the map.
func (h *HashRing) GetNode(key string) string {
	// YOUR CODE HERE
	return ""
}

// GetNodes returns all unique real node addresses currently in the ring.
func (h *HashRing) GetNodes() []string {
	// YOUR CODE HERE
	return nil
}
