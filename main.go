package main

import "api/api"

func main() {
	server := api.New(":3000")
	server.ListenAndServe()
}
