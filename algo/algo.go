package algo

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Algorithm string

const (
	RoundRobin     Algorithm = "round-robin"
	LeastConnections Algorithm = "least-connections"
	IPHash          Algorithm = "ip-hash"
)

type Server struct {
	URL          *url.URL
	ActiveConns  int64
}

type ServerPool struct {
	Servers   []*Server
	current   uint32
	Algorithm Algorithm
}

func (s *ServerPool) GetNextServer(clientIP string) *Server {
	switch s.Algorithm {
	case RoundRobin:
		return s.getNextServerRoundRobin()
	case LeastConnections:
		return s.getNextServerLeastConnections()
	case IPHash:
		return s.getNextServerIPHash(clientIP)
	default:
		return s.getNextServerRoundRobin()
	}
}

func (s *ServerPool) LoadBalance(w http.ResponseWriter, r *http.Request) {
	clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	server := s.GetNextServer(clientIP)
	atomic.AddInt64(&server.ActiveConns, 1)
	defer atomic.AddInt64(&server.ActiveConns, -1)
	proxy := httputil.NewSingleHostReverseProxy(server.URL)
	r.Host = server.URL.Host
	proxy.ServeHTTP(w, r)
}
