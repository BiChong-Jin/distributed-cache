# Distributed Cache

A distributed in-memory key-value cache built from scratch in Go.

## Features

- In-memory key-value store with TTL-based expiration and automatic eviction
- Consistent hashing with virtual nodes for balanced key distribution
- TCP-based client-server communication with a custom wire protocol
- Cluster-aware request routing (local handling or transparent proxying)
- Node discovery with heartbeat-based health checking
- Graceful shutdown via OS signal handling

## Project Structure

```
distributed-cache/
├── main.go              # CLI entry point — start nodes, join cluster
├── cache/               # In-memory store with TTL & eviction
├── consistent/          # Consistent hashing ring for key routing
├── discovery/           # Node registry & heartbeat health checks
├── protocol/            # Wire protocol — request/response serialization
├── server/              # TCP server — routing, local handling, proxying
└── client/              # Client SDK to talk to the cluster
```

## Usage

### Start a cluster

```bash
# Terminal 1 — start the first node
go run main.go -addr :7000

# Terminal 2 — join a second node
go run main.go -addr :7001 -join :7000

# Terminal 3 — join a third node
go run main.go -addr :7002 -join :7000
```

### Build and run

```bash
go build -o distributed-cache
./distributed-cache -addr :7000
```

### Run tests

```bash
go test ./cache/ -race -v
go test ./consistent/ -v
```

## Architecture

1. **Cache Layer** — Each node has a local in-memory store protected by `sync.RWMutex`. Items support TTL, and a background goroutine periodically evicts expired entries.

2. **Consistent Hashing** — Keys are mapped to nodes using a hash ring with virtual nodes. This ensures that adding/removing a node only remaps ~1/N of the keys.

3. **Protocol** — Requests and responses are serialized with `encoding/gob` and sent over raw TCP connections.

4. **Server** — Accepts TCP connections, decodes requests, and routes them. If the current node owns the key, it handles locally; otherwise, it proxies to the correct node.

5. **Client SDK** — Application code uses the client to connect to any node in the cluster. The node handles routing transparently.

6. **Discovery** — Nodes register themselves and send heartbeats. A background goroutine marks unresponsive nodes as suspect or dead.
