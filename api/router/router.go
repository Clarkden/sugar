package router

import (
	"log"
	"net/http"
	"sugar/globals/types"
	"sugar/handlers"
)

func NewRouter(c *types.RouterConfig, h *handlers.Handler) *http.ServeMux {

	if c == nil {
		log.Fatal("Router config not provided")
	}

	router := http.NewServeMux()

	return router
}
