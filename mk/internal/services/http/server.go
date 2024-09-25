package http

import "net/http"

type HHTPServer struct {
	server http.Server
}

func NewHHTPServer() *HHTPServer {
	return &HHTPServer{
		server: http.Server{
			Addr: ":8081",
		},
	}
}

func (s *HHTPServer) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *HHTPServer) Stop() error {
	err := s.server.Close()
	if err != nil {
		return err
	}
	return nil
}
