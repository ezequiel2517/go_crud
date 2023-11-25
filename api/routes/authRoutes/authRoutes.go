package authroutes

import (
	"api/api/controllers/authController"
	"encoding/json"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			authController.SignupHandler(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
			errorMessage := map[string]string{"error": "Método no permitido para la ruta"}
			json.NewEncoder(w).Encode(errorMessage)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			authController.LoginHandler(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
			errorMessage := map[string]string{"error": "Método no permitido para la ruta"}
			json.NewEncoder(w).Encode(errorMessage)
		}
	})
}
