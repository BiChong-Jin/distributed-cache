package server

import (
	"github.com/jin/distributed-cache/cache"
	"github.com/jin/distributed-cache/consistent"
	"github.com/jin/distributed-cache/discovery"
)

// -------- Cache Server (one per node) --------
// Each node runs a Server that:
//   1. Listens on a TCP port for client/peer requests
//   2. Owns a local Cache for the keys hashed to this node
//   3. Uses the HashRing to route requests to the correct node
//   4. Participates in the discovery Registry for cluster membership

// Server is a single node in the distributed cache cluster.
// TODO: Store the listen address, local cache, hash ring, discovery registry,
//       and a net.Listener for accepting TCP connections.
type Server struct {
	Addr     string
	cache    *cache.Cache
	ring     *consistent.HashRing
	registry *discovery.Registry
	// YOUR CODE HERE (net.Listener, shutdown channel, etc.)
}

// NewServer creates a Server but does not start listening yet.
// TODO: Initialize the cache, hash ring, and registry.
func NewServer(addr string) *Server {
	// YOUR CODE HERE
	return nil
}

// Start begins listening on TCP and accepting connections.
// TODO:
//   1. net.Listen("tcp", addr)
//   2. Register self in the registry
//   3. Add self to the hash ring
//   4. In a loop, accept connections and handle each in a goroutine
func (s *Server) Start() error {
	// YOUR CODE HERE
	return nil
}

// Stop gracefully shuts down the server.
// TODO: Close the listener, unregister from discovery, remove from ring.
func (s *Server) Stop() error {
	// YOUR CODE HERE
	return nil
}

// handleConnection reads requests from a TCP connection and sends responses.
// TODO:
//   1. Read bytes from conn → DecodeRequest
//   2. Check if this node owns the key (via hash ring)
//      - If yes: handle locally (get/set/delete on local cache)
//      - If no:  forward the request to the correct node (proxy)
//   3. Encode the Response and write it back
//
// This is the core routing logic of the distributed cache!
func (s *Server) handleConnection( /* net.Conn */ ) {
	// YOUR CODE HERE
}

// handleLocally processes a request against this node's local cache.
// TODO: Switch on CommandType → call cache.Get / cache.Set / cache.Delete.
func (s *Server) handleLocally( /* *protocol.Request */ ) /* *protocol.Response */ {
	// YOUR CODE HERE
}

// forwardToNode sends a request to another node and returns its response.
// TODO: Dial TCP to the target node, send the encoded request, read the response.
func (s *Server) forwardToNode( /* addr string, *protocol.Request */ ) /* *protocol.Response */ {
	// YOUR CODE HERE
}

// JoinCluster adds a known peer node to this server's ring and registry.
// TODO: Add the peer address to the hash ring and register it in discovery.
func (s *Server) JoinCluster(peerAddr string) {
	// YOUR CODE HERE
}
