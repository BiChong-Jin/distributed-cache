package client

import "time"

// -------- Client SDK --------
// This is what application code uses to talk to the cache cluster.
// The client connects to ANY node; that node routes the request to the right place.

// Client is a cache client that connects to a cluster node.
// TODO: Store the server address and optionally a connection pool.
type Client struct {
	// YOUR CODE HERE
}

// NewClient creates a client that talks to the cache cluster via the given node address.
func NewClient(addr string) *Client {
	// YOUR CODE HERE
	return nil
}

// Close cleans up any open connections.
func (c *Client) Close() error {
	// YOUR CODE HERE
	return nil
}

// Set stores a key-value pair with the given TTL.
// TODO: Build a protocol.Request with CmdSet, encode it, send over TCP, read response.
func (c *Client) Set(key string, value []byte, ttl time.Duration) error {
	// YOUR CODE HERE
	return nil
}

// Get retrieves a value by key.
// TODO: Build a protocol.Request with CmdGet, send it, decode response.
func (c *Client) Get(key string) ([]byte, error) {
	// YOUR CODE HERE
	return nil, nil
}

// Delete removes a key.
func (c *Client) Delete(key string) error {
	// YOUR CODE HERE
	return nil
}

// Keys returns all keys in the cluster (queries the connected node).
func (c *Client) Keys() ([]string, error) {
	// YOUR CODE HERE
	return nil, nil
}

// Ping checks whether the connected node is alive.
func (c *Client) Ping() error {
	// YOUR CODE HERE
	return nil
}

// sendRequest is a helper that handles the TCP send/receive cycle.
// TODO: Dial (or reuse connection), write encoded request, read and decode response.
func (c *Client) sendRequest( /* *protocol.Request */ ) /* (*protocol.Response, error) */ {
	// YOUR CODE HERE
}
