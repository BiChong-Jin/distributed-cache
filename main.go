package main

import (
	"flag"
	"fmt"
	"log"
)

// -------- Entry Point --------
// This starts a cache node. Usage:
//   go run main.go -addr :7000
//   go run main.go -addr :7001 -join :7000
//   go run main.go -addr :7002 -join :7000

func main() {
	addr := flag.String("addr", ":7000", "listen address for this node")
	join := flag.String("join", "", "address of an existing node to join the cluster")
	flag.Parse()

	fmt.Printf("Starting cache node on %s\n", *addr)
	if *join != "" {
		fmt.Printf("Joining cluster via %s\n", *join)
	}

	// TODO:
	// 1. Create a new Server with the given addr
	// 2. If -join is provided, call server.JoinCluster(joinAddr)
	// 3. Call server.Start() (this blocks)
	// 4. Handle OS signals (SIGINT, SIGTERM) for graceful shutdown → server.Stop()

	_ = addr
	_ = join
	log.Fatal("Not implemented yet — start coding!")
}
