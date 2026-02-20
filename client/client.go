package client

import (
	"encoding/json"
	"net"
	"time"

	"github.com/BiChong-Jin/distributed-cache/protocol"
)

// -------- Client SDK --------
// This is what application code uses to talk to the cache cluster.
// The client connects to ANY node; that node routes the request to the right place.

// Client is a cache client that connects to a cluster node.
type Client struct {
	Addr string
}

// NewClient creates a client that talks to the cache cluster via the given node address.
func NewClient(addr string) *Client {
	return &Client{Addr: addr}
}

// Close cleans up any open connections.
func (c *Client) Close() error {
	return nil
}

// Set stores a key-value pair with the given TTL.
func (c *Client) Set(key string, value []byte, ttl time.Duration) error {
	req := &protocol.Request{
		CommandType: protocol.CmdSet,
		Key:         key,
		Value:       value,
		TTL:         ttl,
	}

	_, err := c.sendRequest(req)
	if err != nil {
		return err
	}

	return nil
}

// Get retrieves a value by key.
func (c *Client) Get(key string) ([]byte, error) {
	req := &protocol.Request{
		CommandType: protocol.CmdGet,
		Key:         key,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

// Delete removes a key.
func (c *Client) Delete(key string) error {
	req := &protocol.Request{
		CommandType: protocol.CmdDelete,
		Key:         key,
	}

	_, err := c.sendRequest(req)
	if err != nil {
		return err
	}
	return nil
}

// Keys returns all keys in the cluster (queries the connected node).
func (c *Client) Keys() ([]string, error) {
	req := &protocol.Request{
		CommandType: protocol.CmdKeys,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var keys []string
	err = json.Unmarshal(resp.Value, &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

// Ping checks whether the connected node is alive.
func (c *Client) Ping() error {
	req := &protocol.Request{
		CommandType: protocol.CmdPing,
	}

	_, err := c.sendRequest(req)
	if err != nil {
		return err
	}
	return nil
}

// sendRequest is a helper that handles the TCP send/receive cycle.
func (c *Client) sendRequest(req *protocol.Request) (*protocol.Response, error) {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		return &protocol.Response{StatusCode: protocol.StatusError, ErrorMessage: "Cannot connect to the cluster."}, err
	}
	defer conn.Close()

	reqBytes, err := req.Encode()
	if err != nil {
		return &protocol.Response{StatusCode: protocol.StatusError, ErrorMessage: "Bad request."}, err
	}

	conn.Write(reqBytes)
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return &protocol.Response{StatusCode: protocol.StatusError, ErrorMessage: "Cannot get response."}, err
	}
	resp := buf[:n]

	return protocol.DecodeResponse(resp)
}
