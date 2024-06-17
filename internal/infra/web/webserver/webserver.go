package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	HTTP_POST = "POST"
	HTTP_GET  = "GET"
)

type Handler struct {
	Func   http.HandlerFunc
	Method string
	Path   string
}

type WebServer struct {
	Router        chi.Router
	Handlers      []Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []Handler{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path, method string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, Handler{
		Func:   handler,
		Method: method,
		Path:   path,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		if handler.Method == HTTP_POST {
			s.Router.Post(handler.Path, handler.Func)
		} else {
			s.Router.Get(handler.Path, handler.Func)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
