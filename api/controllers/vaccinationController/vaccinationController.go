package vaccinationController

import (
	"api/api/controllers/drugController"
	"api/api/models"
	"api/db"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
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
	INSERT INTO vaccinations (name, drug_id, dose, fecha)
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

func GetVaccinations(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()
	sqlStatement := `
	SELECT id, name, drug_id, dose, fecha
	FROM vaccinations
	`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al obtener vacunaciones de la base de datos."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	defer rows.Close()

	var vaccinations []models.Vaccination
	for rows.Next() {
		var vaccination models.Vaccination
		err := rows.Scan(&vaccination.ID, &vaccination.Name, &vaccination.DrugId, &vaccination.DrugId, &vaccination.Fecha)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorMessage := map[string]string{"error": "Error al obtener drogas de la base de datos."}
			json.NewEncoder(w).Encode(errorMessage)
			return
		}
		vaccinations = append(vaccinations, vaccination)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vaccinations)
}

func DeleteVaccination(w http.ResponseWriter, r *http.Request) {
	vId := r.URL.Path[len("/vaccinations/"):]
	if vId == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "ID de vacunación no proporcionado."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	id, err := strconv.Atoi(vId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "ID de vacunación no es un entero."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	db := db.GetConnection()
	sqlStatement := `
	DELETE FROM vaccinations
	WHERE id = $1`

	_, err = db.Exec(sqlStatement, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := map[string]string{"error": "Error al eliminar vacunación de la base de datos."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	errorMessage := map[string]string{"ok": "Vacunación eliminada correctamente."}
	json.NewEncoder(w).Encode(errorMessage)
}
