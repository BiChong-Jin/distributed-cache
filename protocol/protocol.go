package protocol

import "time"

// -------- Wire Protocol --------
// Defines how nodes talk to each other over TCP.
// You'll serialize these structs into bytes to send over the network.

// CommandType identifies what operation a request is asking for.
type CommandType byte

const (
	CmdGet    CommandType = iota + 1 // Retrieve a value
	CmdSet                           // Store a value
	CmdDelete                        // Remove a value
	CmdPing                          // Health check
	CmdKeys                          // List all keys
)

// StatusCode indicates success or failure in a response.
type StatusCode byte

const (
	StatusOK       StatusCode = iota + 1
	StatusNotFound
	StatusError
)

// Request is the message a client sends to a cache node.
// TODO: Include the command type, key, value (for Set), and TTL.
type Request struct {
  commandType CommandType
  key int
  value string
  ttl time.Duration
}

// Response is the message a cache node sends back to a client.
// TODO: Include the status code, value (for Get), and an error message if any.
type Response struct {
  statusCode StatusCode
  value string
  errorMessage string
}

// -------- Serialization --------
// You need to convert Request/Response to []byte and back.
// Choose one approach:
//   Option A (easier): Use encoding/gob or encoding/json
//   Option B (learn more): Build a custom binary protocol
//     e.g. [1 byte cmd][2 bytes key len][key][4 bytes val len][value][8 bytes ttl]

// Encode serializes a Request into bytes for sending over the network.
// TODO: Use your chosen encoding strategy.
func (r *Request) Encode() ([]byte, error) {
	// YOUR CODE HERE
	return nil, nil
}

// DecodeRequest deserializes bytes back into a Request.
func DecodeRequest(data []byte) (*Request, error) {
	// YOUR CODE HERE
	return nil, nil
}

// Encode serializes a Response into bytes.
func (r *Response) Encode() ([]byte, error) {
	// YOUR CODE HERE
	return nil, nil
}

// DecodeResponse deserializes bytes back into a Response.
func DecodeResponse(data []byte) (*Response, error) {
	// YOUR CODE HERE
	return nil, nil
}

// Ensure the compiler knows we use time.Duration (suppress unused import).
var _ time.Duration
