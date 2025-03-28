package main

import (
	"auth/config"
	"auth/handlers"
	"auth/utils"
	"fmt"
	"net/http"
)

func main() {
	config.InitGlobalEnv()

	mux := http.NewServeMux()
	port := utils.GetEnv("AUTH_PORT")

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.Signup(w, r)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("Server running on http://localhost:%s", port)
	http.ListenAndServe(":"+port, mux)
}
