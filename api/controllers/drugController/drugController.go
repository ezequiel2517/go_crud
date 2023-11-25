package drugController

import (
	"api/api/models"
	"api/db"
	"encoding/json"
	"net/http"
	"strconv"
)

func InsertDrug(w http.ResponseWriter, r *http.Request) {
	var drug models.Drug
	err := json.NewDecoder(r.Body).Decode(&drug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al convertir body."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	db := db.GetConnection()
	sqlStatement := `
	INSERT INTO drugs (name, approved, min_dose, max_dose, available_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	_, err = db.Exec(sqlStatement, drug.Name, drug.Approved, drug.MinDose, drug.MaxDose, drug.AvailableAt)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al insertar droga en la base de datos."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.WriteHeader(http.StatusCreated)
	errorMessage := map[string]string{"ok": "Droga creada exitosamente."}
	json.NewEncoder(w).Encode(errorMessage)
}

func UpdateDrug(w http.ResponseWriter, r *http.Request) {
	drugId := r.URL.Path[len("/drugs/"):]
	if drugId == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "ID de drug no proporcionado."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	id, err := strconv.Atoi(drugId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "ID de drug no es un entero."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	var drug models.Drug
	err = json.NewDecoder(r.Body).Decode(&drug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al convertir body."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	db := db.GetConnection()
	sqlStatement := `
	UPDATE drugs
	SET name = $1, approved= $2 , 
	min_dose = $3, max_dose= $4,
	available_at = $5 
	WHERE id = $6`

	_, err = db.Exec(sqlStatement, drug.Name, drug.Approved, drug.MinDose, drug.MaxDose, drug.AvailableAt, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := map[string]string{"error": "Error al actualizar droga en la base de datos."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	errorMessage := map[string]string{"ok": "Droga actualizada correctamente."}
	json.NewEncoder(w).Encode(errorMessage)
}

func GetDrugs(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()
	sqlStatement := `
	SELECT id, name, approved, min_dose, max_dose, available_at
	FROM drugs
	`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al obtener drogas de la base de datos."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	defer rows.Close()

	var drugs []models.Drug
	for rows.Next() {
		var drug models.Drug
		err := rows.Scan(&drug.ID, &drug.Name, &drug.Approved, &drug.MinDose, &drug.MaxDose, &drug.AvailableAt)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorMessage := map[string]string{"error": "Error al obtener drogas de la base de datos."}
			json.NewEncoder(w).Encode(errorMessage)
			return
		}
		drugs = append(drugs, drug)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drugs)
}

func GetDrug(id int) (models.Drug, error) {
	db := db.GetConnection()
	sqlStatement := `
	SELECT id, name, approved, min_dose, max_dose, available_at
	FROM drugs
	WHERE id = $1
	`

	row := db.QueryRow(sqlStatement, id)
	var drug models.Drug
	err := row.Scan(&drug.ID, &drug.Name, &drug.Approved, &drug.MinDose, &drug.MaxDose, &drug.AvailableAt)
	if err != nil {
		return models.Drug{}, err
	}
	return drug, nil
}

func DeleteDrug(w http.ResponseWriter, r *http.Request) {
	drugId := r.URL.Path[len("/drugs/"):]
	if drugId == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "ID de drug no proporcionado."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	id, err := strconv.Atoi(drugId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "ID de drug no es un entero."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	db := db.GetConnection()
	sqlStatement := `
	DELETE FROM drugs
	WHERE id = $1`

	_, err = db.Exec(sqlStatement, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := map[string]string{"error": "Error al eliminar droga de la base de datos."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	errorMessage := map[string]string{"ok": "Droga eliminada correctamente."}
	json.NewEncoder(w).Encode(errorMessage)
}
