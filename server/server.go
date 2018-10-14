package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/teran/microgpio/drivers"
)

// Server type
type Server struct {
	driver drivers.Driver
	e      *echo.Echo
}

// New returns new instance of Server
func New(driver drivers.Driver) *Server {
	s := &Server{
		driver: driver,
		e:      echo.New(),
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
	s.e.POST("/gpio/:id/high", s.high)
	s.e.POST("/gpio/:id/low", s.low)
	s.e.POST("/gpio/:id/output", s.output)
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
