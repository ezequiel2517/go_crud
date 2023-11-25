package vaccionationRoutes

import (
	"api/api/controllers/vaccinationController"
	"encoding/json"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/vaccinations", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			vaccinationController.InsertVaccination(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
			errorMessage := map[string]string{"error": "Método no permitido para la ruta"}
			json.NewEncoder(w).Encode(errorMessage)
		}
	})
}
