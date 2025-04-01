package main

import (
	"fmt"
	"net/http"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // Set Content-Type
	w.WriteHeader(http.StatusOK)                                // Set status code 200
	w.Write([]byte("OK"))                                       // Write "OK" as the response body
}

func main() {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/healthz", readinessHandler)

	serverMux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))

	server := &http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server error: ", err)
	}
}
