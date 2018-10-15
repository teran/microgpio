package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server type
type Server struct {
	e *echo.Echo
}

// New returns new instance of Server
func New() *Server {
	s := &Server{
		e: echo.New(),
	}

	// Middleware
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.populateRoutes()

	return s
}

func (s *Server) populateRoutes() {
	s.e.GET("/", s.index)
	s.e.GET("/ping", s.ping)
	s.e.POST("/gpio/:id/export", s.export)
	s.e.POST("/gpio/:id/unexport", s.unexport)
	s.e.POST("/gpio/:id/high", s.high)
	s.e.POST("/gpio/:id/low", s.low)
	s.e.GET("/gpio/:id/mode", s.mode)
	s.e.POST("/gpio/:id/input", s.input)
	s.e.POST("/gpio/:id/output", s.output)
	s.e.GET("/gpio/:id/value", s.value)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.e.ServeHTTP(w, r)
}

// ListenAndServe starts http server
func (s *Server) ListenAndServe(addr string) error {
	srv := http.Server{
		Addr:    addr,
		Handler: s,
	}

	return srv.ListenAndServe()
}
