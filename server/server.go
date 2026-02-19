package server

import (
	"encoding/json"
	"net"
	"time"

	"github.com/BiChong-Jin/distributed-cache/cache"
	"github.com/BiChong-Jin/distributed-cache/consistent"
	"github.com/BiChong-Jin/distributed-cache/discovery"
	"github.com/BiChong-Jin/distributed-cache/protocol"
)

// -------- Cache Server (one per node) --------
// Each node runs a Server that:
//   1. Listens on a TCP port for client/peer requests
//   2. Owns a local Cache for the keys hashed to this node
//   3. Uses the HashRing to route requests to the correct node
//   4. Participates in the discovery Registry for cluster membership

// Server is a single node in the distributed cache cluster.
type Server struct {
	Addr     string
	cache    *cache.Cache
	ring     *consistent.HashRing
	registry *discovery.Registry
	listener net.Listener
}

// NewServer creates a Server but does not start listening yet.
func NewServer(addr string) *Server {
	return &Server{
		Addr:     addr,
		cache:    cache.NewCache(5 * time.Second),
		ring:     consistent.NewHashRing(150),
		registry: discovery.NewRegistry(10 * time.Second),
	}
}

// Start begins listening on TCP and accepting connections.
func (s *Server) Start() error {
	addr := s.Addr
	hr := s.ring
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.listener = listener
	s.registry.Register(addr)
	hr.AddNode(addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go s.handleConnection(conn)
	}
	return nil
}

// Stop gracefully shuts down the server.
func (s *Server) Stop() error {
	err := s.listener.Close()
	if err != nil {
		return err
	}

	s.registry.Unregister(s.Addr)
	s.ring.RemoveNode(s.Addr)

	return nil
}

// handleConnection reads requests from a TCP connection and sends responses.
//  1. Read bytes from conn → DecodeRequest
//  2. Check if this node owns the key (via hash ring)
//     - If yes: handle locally (get/set/delete on local cache)
//     - If no:  forward the request to the correct node (proxy)
//  3. Encode the Response and write it back
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	data := buf[:n]
	req, err := protocol.DecodeRequest(data)
	if err != nil {
		return
	}

	k := req.Key
	no := s.ring.GetNode(k)
	var res *protocol.Response

	if s.Addr == no {
		res = s.handleLocally(req)
	} else {
		res = s.forwardToNode(no, req)
	}

	respBytes, err := res.Encode()
	if err != nil {
		return
	}

	conn.Write(respBytes)
}

// handleLocally processes a request against this node's local cache.
func (s *Server) handleLocally(req *protocol.Request) *protocol.Response {
	switch req.CommandType {
	case protocol.CmdGet:
		val, ok := s.cache.Get(req.Key)
		if !ok {
			return &protocol.Response{StatusCode: protocol.StatusNotFound}
		}
		return &protocol.Response{StatusCode: protocol.StatusOK, Value: val}

	case protocol.CmdSet:
		s.cache.Set(req.Key, req.Value, req.TTL)
		return &protocol.Response{StatusCode: protocol.StatusOK}

	case protocol.CmdDelete:
		s.cache.Delete(req.Key)
		return &protocol.Response{StatusCode: protocol.StatusOK}

	case protocol.CmdPing:
		return &protocol.Response{StatusCode: protocol.StatusOK}

	case protocol.CmdKeys:
		keys := s.cache.Keys()
		data, err := json.Marshal(keys)
		if err != nil {
			return &protocol.Response{StatusCode: protocol.StatusError, ErrorMessage: "Failed to get keys."}
		}
		return &protocol.Response{StatusCode: protocol.StatusOK, Value: data}

	default:
		return &protocol.Response{StatusCode: protocol.StatusError, ErrorMessage: "Unknown CommandType."}
	}
}

// forwardToNode sends a request to another node and returns its response.
func (s *Server) forwardToNode(addr string, req *protocol.Request) *protocol.Response {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return &protocol.Response{StatusCode: protocol.StatusError, ErrorMessage: err.Error()}
	}
	defer conn.Close()

	reqBytes, err := req.Encode()
	if err != nil {
		return &protocol.Response{StatusCode: protocol.StatusError, ErrorMessage: err.Error()}
	}
	conn.Write(reqBytes)

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return &protocol.Response{StatusCode: protocol.StatusError, ErrorMessage: err.Error()}
	}

	res, err := protocol.DecodeResponse(buf[:n])
	if err != nil {
		return &protocol.Response{StatusCode: protocol.StatusError, ErrorMessage: err.Error()}
	}

	return res
}

// JoinCluster adds a known peer node to this server's ring and registry.
func (s *Server) JoinCluster(peerAddr string) {
	s.ring.AddNode(peerAddr)
	s.registry.Register(peerAddr)
}
