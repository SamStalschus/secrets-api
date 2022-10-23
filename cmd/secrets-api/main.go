package main

import (
	"fmt"
	"log"
	"net/http"
	"secrets-api/infra/env"
)

func main() {
	port := env.GetString("PORT", "8080")

	if err := run(port); err != nil {
		log.Fatal(fmt.Sprintf("Error to start server on port: %s - Erro: %s ", port, err))
	}
}

func run(port string) error {
	log.Println("Listening on port", port)
	handler := http.HandlerFunc(serve)
	return http.ListenAndServe(":"+port, handler)
}
