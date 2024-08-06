package algo

import (
	"hash/fnv"
)

func (s *ServerPool) getNextServerIPHash(clientIP string) *Server {
	hash := fnv.New32a()
	hash.Write([]byte(clientIP))
	index := hash.Sum32() % uint32(len(s.Servers))
	return s.Servers[index]
}
