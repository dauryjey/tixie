package main

import (
	"auth/config"
	"fmt"
	"net/http"
	"os"
)

func main() {
	config.InitGlobalEnv()

	mux := http.NewServeMux()
	port := os.Getenv("AUTH_PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on http://localhost:%s", port)
	http.ListenAndServe(port, mux)
}
