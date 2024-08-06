package algo

import "sync/atomic"

func (s *ServerPool) getNextServerLeastConnections() *Server {
	var leastConnServer *Server
	for _, server := range s.Servers {
		if leastConnServer == nil || atomic.LoadInt64(&server.ActiveConns) < atomic.LoadInt64(&leastConnServer.ActiveConns) {
			leastConnServer = server
		}
	}
	return leastConnServer
}
