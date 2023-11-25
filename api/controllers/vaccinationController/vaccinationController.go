package vaccinationController

import (
	"api/api/controllers/drugController"
	"api/api/models"
	"api/db"
	"database/sql"
	"encoding/json"
	"net/http"
)

func InsertVaccination(w http.ResponseWriter, r *http.Request) {
	var vaccination models.Vaccination
	err := json.NewDecoder(r.Body).Decode(&vaccination)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al convertir body."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	droga, err := drugController.GetDrug(vaccination.DrugId)

	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "ID de droga no existe."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al validar ID de droga."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if droga.AvailableAt.After(vaccination.Fecha) {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Fecha de vacunación previa a fecha de disponibilidad de droga."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	db := db.GetConnection()
	sqlStatement := `
	INSERT INTO vaccination (name, drug_id, dose, fecha)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	_, err = db.Exec(sqlStatement, vaccination.Name, vaccination.DrugId, vaccination.Dose, vaccination.Fecha)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al insertar vacunación en la base de datos."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.WriteHeader(http.StatusCreated)
	errorMessage := map[string]string{"ok": "Vacunación creada exitosamente."}
	json.NewEncoder(w).Encode(errorMessage)
}
