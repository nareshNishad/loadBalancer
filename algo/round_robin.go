package algo

import "sync/atomic"

func (s *ServerPool) getNextServerRoundRobin() *Server {
	next := atomic.AddUint32(&s.current, 1)
	return s.Servers[(int(next)-1)%len(s.Servers)]
}
