package main

import (
	"api/api"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func cargarEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR AL CARGAR ARCHIVO .env")
	}
}

func main() {
	cargarEnv()
	server := api.New(os.Getenv("PUERTO"))
	server.ListenAndServe()
}
