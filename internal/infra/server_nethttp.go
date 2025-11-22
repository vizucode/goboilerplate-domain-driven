package infra

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type server struct {
	mux         *http.ServeMux
	Middlewares []Middleware
}

func NewNetHttpServer() *server {
	return &server{
		mux: http.NewServeMux(),
	}
}

func (s *server) Use(mw Middleware) {
	s.Middlewares = append(s.Middlewares, mw)
}

func (s *server) NetHttpListen(addr string) {
	http.ListenAndServe(addr, s.mux)
}

func chain(handler http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		handler = mws[i](handler)
	}
	return handler
}

func (s *server) GET(path string, handler http.HandlerFunc, routeMiddlewares ...Middleware) {
	finalHandler := http.Handler(handler)

	if len(routeMiddlewares) > 0 {
		finalHandler = chain(finalHandler, routeMiddlewares...)
	}

	if len(s.Middlewares) > 0 {
		finalHandler = chain(finalHandler, s.Middlewares...)
	}

	finalHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		finalHandler.ServeHTTP(w, r)
	})

	s.mux.Handle(path, finalHandler)
}

func (s *server) POST(path string, handler http.HandlerFunc, routeMiddlewares ...Middleware) {
	finalHandler := http.Handler(handler)

	if len(routeMiddlewares) > 0 {
		finalHandler = chain(finalHandler, routeMiddlewares...)
	}

	if len(s.Middlewares) > 0 {
		finalHandler = chain(finalHandler, s.Middlewares...)
	}

	finalHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		finalHandler.ServeHTTP(w, r)
	})

	s.mux.Handle(path, finalHandler)
}

func (s *server) PUT(path string, handler http.HandlerFunc, routeMiddlewares ...Middleware) {
	finalHandler := http.Handler(handler)

	if len(routeMiddlewares) > 0 {
		finalHandler = chain(finalHandler, routeMiddlewares...)
	}

	if len(s.Middlewares) > 0 {
		finalHandler = chain(finalHandler, s.Middlewares...)
	}

	finalHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		finalHandler.ServeHTTP(w, r)
	})

	s.mux.Handle(path, finalHandler)
}

func (s *server) DELETE(path string, handler http.HandlerFunc, routeMiddlewares ...Middleware) {
	finalHandler := http.Handler(handler)

	if len(routeMiddlewares) > 0 {
		finalHandler = chain(finalHandler, routeMiddlewares...)
	}

	if len(s.Middlewares) > 0 {
		finalHandler = chain(finalHandler, s.Middlewares...)
	}

	finalHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		finalHandler.ServeHTTP(w, r)
	})

	s.mux.Handle(path, finalHandler)
}
