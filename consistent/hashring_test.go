package consistent

import "testing"

func TestAddAndGetNode(t *testing.T) {
	// TODO: Create a ring, add 3 nodes, verify GetNode returns one of them for any key.
}

func TestDistribution(t *testing.T) {
	// TODO: Add 3 nodes, hash 10000 random keys, count how many land on each node.
	// Verify no node gets less than 20% or more than 45% (rough balance check).
}

func TestRemoveNode(t *testing.T) {
	// TODO: Add 3 nodes, record which node owns "test-key".
	// Remove a DIFFERENT node, verify "test-key" still maps to the same node.
}

func TestConsistencyAfterAddition(t *testing.T) {
	// TODO: Add 2 nodes, record mapping for 1000 keys.
	// Add a 3rd node, verify that the majority (>60%) of keys still map to the same node.
}

func TestEmptyRing(t *testing.T) {
	// TODO: GetNode on an empty ring should return "".
}
