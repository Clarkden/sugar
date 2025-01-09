package router

import (
	"log"
	"net/http"
	"sugar/globals/types"
	"sugar/handlers"
	"sugar/middleware"
)

func NewRouter(c *types.RouterConfig, h *handlers.Handler) *http.ServeMux {

	if c == nil {
		log.Fatal("Router config not provided")
	}

	router := http.NewServeMux()

	v1Group := func(pattern string, handler http.HandlerFunc, method string, noAuth ...bool) {
		var finalHandler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			handler.ServeHTTP(w, r)
		})

		if len(noAuth) == 0 || !noAuth[0] {
			finalHandler = middleware.AuthMiddleware(finalHandler)
		}

		router.Handle(method+" /v1"+pattern, finalHandler)
	}

	authGroup := func(pattern string, handler http.HandlerFunc, method string, noAuth ...bool) {
		v1Group("/auth"+pattern, handler, method, noAuth...)
	}

	authGroup("/email/register", h.HandleEmailRegister(), http.MethodPost)
	authGroup("/email/login", h.HandleEmailLogin(), http.MethodPost)

	return router
}
