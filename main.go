package main

import (
	"fmt"
	"loadbalancer/algo"
	"loadbalancer/server"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <algorithm>")
		fmt.Println("Available algorithms: round-robin, least-connections, ip-hash")
		return
	}

	algorithm := algo.Algorithm(os.Args[1])
	switch algorithm {
	case algo.RoundRobin, algo.LeastConnections, algo.IPHash:
	default:
		fmt.Println("Invalid algorithm. Available algorithms: round-robin, least-connections, ip-hash")
		return
	}

	servers := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	var serverPool algo.ServerPool
	serverPool.Algorithm = algorithm

	for _, server := range servers {
		url, err := url.Parse(server)
		if err != nil {
			log.Fatalf("Error parsing server URL: %v", err)
		}
		serverPool.Servers = append(serverPool.Servers, &algo.Server{URL: url})
	}


	go server.RunServers()

	http.HandleFunc("/", serverPool.LoadBalance)
	fmt.Printf("Load Balancer started at :8080 using %s algorithm\n", algorithm)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
