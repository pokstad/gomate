package http

import (
	"context"
	"net"
	"net/http"
)

type Server struct {
	s *http.Server
}

func NewServer() Server {
	s := Server{
		s: &http.Server{},
	}
	return s
}

func (s Server) Start(ctx context.Context) error {
	l, err := net.ListenUnix("unix", &net.UnixAddr{
		Name: "/tmp/gomate.sock",
		Net:  "unix",
	})
	if err != nil {
		return err
	}

	return s.s.Serve(l)
}

func outline(rw http.ResponseWriter, req *http.Request) {

}

func init() {
	http.HandleFunc("/outline", outline)
}
