package api

import (
	"api/api/routes"
	"net/http"
)

func New(addr string) *http.Server {
	routes.InitRoutes()
	return &http.Server{
		Addr: addr,
	}
}
