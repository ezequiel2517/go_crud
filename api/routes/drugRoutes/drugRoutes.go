package drugRoutes

import (
	"api/api/controllers/drugController"
	"encoding/json"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/drugs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			drugController.InsertDrug(w, r)
		case http.MethodGet:
			drugController.GetDrugs(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
			errorMessage := map[string]string{"error": "Método no permitido para la ruta"}
			json.NewEncoder(w).Encode(errorMessage)
		}
	})

	http.HandleFunc("/drugs/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			drugController.UpdateDrug(w, r)
		case http.MethodDelete:
			drugController.DeleteDrug(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
			errorMessage := map[string]string{"error": "Método no permitido para la ruta"}
			json.NewEncoder(w).Encode(errorMessage)
		}
	})
}
