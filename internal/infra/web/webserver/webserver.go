package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      []HandlerMethod
	WebServerPort string
}

type HandlerMethod struct {
	Method string
	Path   string
	Func   http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]HandlerMethod, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, HandlerMethod{method, path, handler})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, h := range s.Handlers {
		s.Router.Method(h.Method, h.Path, h.Func)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
