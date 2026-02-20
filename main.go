package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/BiChong-Jin/distributed-cache/server"
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

	s := server.NewServer(*addr)
	if *join != "" {
		s.JoinCluster(*join)
	}
	go s.Start()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	s.Stop()
}
