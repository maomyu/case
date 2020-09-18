package server

import (
	"net/http"

	_ "net/http/pprof"
)

type Server struct {
	Server        *http.Server
}

func (s *Server)GetServerHttp(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("111"))
}