package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method   string
	Path     string
	Handlers []gin.HandlerFunc
}

type HTTPServer struct {
	routes []Route
	port   string
}

func NewHTTPServer(port string) *HTTPServer {
	return &HTTPServer{
		port:   port,
		routes: []Route{},
	}
}

func (s *HTTPServer) AddRoute(method string, path string, handlers ...gin.HandlerFunc) {
	s.routes = append(s.routes, Route{Method: method, Path: path, Handlers: handlers})
}

func (s *HTTPServer) Run() error {
	r := gin.Default()

	for _, rt := range s.routes {
		r.Handle(rt.Method, rt.Path, rt.Handlers...)
	}

	log.Printf("HTTP server listening on %s", s.port)
	return http.ListenAndServe(s.port, r)
}
