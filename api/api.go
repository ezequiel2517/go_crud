package api

import (
	authroutes "api/api/routes/authRoutes"
	"api/api/routes/drugRoutes"
	vaccionationRoutes "api/api/routes/vaccinationRoutes"
	"net/http"
)

func New(addr string) *http.Server {
	authroutes.InitRoutes()
	drugRoutes.InitRoutes()
	vaccionationRoutes.InitRoutes()
	return &http.Server{
		Addr: addr,
	}
}
