package main

import (
	"auth/config"
	"auth/utils"
	"fmt"
	"net/http"
)

func main() {
	config.InitGlobalEnv()

	mux := http.NewServeMux()
	port := utils.GetEnv("PORT")

	fmt.Printf("Server running on http://localhost:%s", port)
	http.ListenAndServe(port, mux)
}
