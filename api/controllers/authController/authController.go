package authController

import (
	"api/api/authentication"
	usercontroller "api/api/controllers/userController"
	"api/api/models"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin models.User
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al convertir body."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	user, _ := usercontroller.GetUser(userLogin.Email)
	if !authentication.CheckPassword(userLogin.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		errorMessage := map[string]string{"error": "Credenciales invalidas."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	tokenStr, err := authentication.GenerarToken(userLogin.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := map[string]string{"error": "Error al generar el JWT."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al convertir body."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	userExiste, _ := usercontroller.GetUser(newUser.Email)
	if userExiste != (models.User{}) {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Usuario ya existe."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	hashedPassword, err := authentication.HashPassword(newUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := map[string]string{"error": "Error al aplicar hash a password."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	err = usercontroller.InsertUser(models.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: hashedPassword,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Error al registrar el usuario en la base de datos."}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.WriteHeader(http.StatusCreated)
	errorMessage := map[string]string{"ok": "Usuario registrado con exito."}
	json.NewEncoder(w).Encode(errorMessage)
}
