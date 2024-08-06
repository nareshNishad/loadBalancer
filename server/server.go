package server

import (
	"log"
	"net/http"
	"sync"
)

func StartServer(port string, handler http.HandlerFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())
}

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Response from server 1"))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Response from server 2"))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Response from server 3"))
}

func RunServers() {
	var wg sync.WaitGroup
	wg.Add(3)

	go StartServer(":8081", handler1, &wg)
	go StartServer(":8082", handler2, &wg)
	go StartServer(":8083", handler3, &wg)

	wg.Wait()
}
