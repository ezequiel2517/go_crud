package api

import (
	"api/api/routes/drugRoutes"
	vaccionationRoutes "api/api/routes/vaccinationRoutes"
	"net/http"
)

func New(addr string) *http.Server {
	drugRoutes.InitRoutes()
	vaccionationRoutes.InitRoutes()
	return &http.Server{
		Addr: addr,
	}
}
