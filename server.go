package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server 8080 portunda çalışıyor...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server hatası: %s\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selam! Go HTTP Server'a hoş geldin. Saat: %s", time.Now().Format("15:04:05"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
