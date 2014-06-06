package commandserver

import (
	"log"
	"net"
	"net/textproto"
)

type Server struct {
	address string
	mux     *ServeMux
}

func NewServer(address string) *Server {
	return &Server{
		address: address,
	}
}

func (s *Server) ListenAndServe(mux *ServeMux) error {
	s.mux = mux

	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s\n", err)
		}
		go s.serve(conn)
	}
}

func (s *Server) serve(conn net.Conn) {
	c := textproto.NewConn(conn)

	// state controlled processing here

	for {
		cr, err := NewCommandRequest(c)
		if err != nil {
			log.Println(err)
			continue
		}

		// start of a muxer
		// s.handler[verb](c)

		sr := &StatusResponse{}

		if s.mux == nil {
			panic("awp: no servemux specified")
		}
		s.mux.ServeAWP(cr, sr)

		c.PrintfLine("%d %s\r\n", sr.Code, sr.Status)

	}
}
