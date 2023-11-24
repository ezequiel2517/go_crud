package routes

import (
	"api/api/handlers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/drugs", handlers.InsertarDrug)
}
