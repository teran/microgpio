package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/teran/microgpio/controller"
)

// Server type
type Server struct {
	e *echo.Echo
	c *controller.Controller
}

// New returns new instance of Server
func New(c *controller.Controller) *Server {
	s := &Server{
		e: echo.New(),
		c: c,
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

	s.e.POST("/pin/:name/on", s.on)
	s.e.POST("/pin/:name/off", s.off)
	s.e.GET("/pin/:name", s.status)
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
