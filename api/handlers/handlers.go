package handlers

import (
	"api/api/models"
	"api/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}

func InsertarDrug(w http.ResponseWriter, r *http.Request) {
	db := repository.GetConnection()

	var drug models.Drug
	err := json.NewDecoder(r.Body).Decode(&drug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error al convertir body. Revisar solicitud.")
		return
	}

	err = models.InsertarDrug(db, drug.Name, drug.Approved, drug.MinDose, drug.MaxDose, drug.AvailableAt)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al insertar droga en la base de datos.")
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Droga creada exitosamente")

	db.Close()
}
