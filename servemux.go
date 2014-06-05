package commandserver

import (
	"sync"
)

type HandlerFunc func(c *CommandRequest, r *StatusResponse)

type ServeMux struct {
	mu sync.RWMutex
	m  map[string]HandlerFunc
}

func NewServeMux() *ServeMux {
	return &ServeMux{m: make(map[string]HandlerFunc)}
}

func (mux *ServeMux) Handle(pattern string, handler HandlerFunc) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("cmdsrv: invalid pattern " + pattern)
	}

	if handler == nil {
		panic("cmdsrv: nil handler")
	}

	mux.m[pattern] = handler
}

func (mux *ServeMux) ServeAWP(c *CommandRequest, s *StatusResponse) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()
	if h, ok := mux.m[c.Method]; ok {
		h(c, s)
	} else {
		// Bad Method/404
		s.Code = 404
		s.Status = "Unrecognised command"
	}
}
