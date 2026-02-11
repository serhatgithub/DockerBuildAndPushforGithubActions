package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/version", versionHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("🚀 Server v2.0 başlatılıyor... Port: 8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server hatası: %s\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Bilinmiyor"
	}

	message := fmt.Sprintf("Merhaba! Burası v2 Sürümü.\n\nÇalışan Sunucu (Hostname): %s\nZaman: %s", hostname, time.Now().Format("15:04:05"))

	fmt.Fprintf(w, message)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK - Sistem Ayakta"))
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v2.0.0"))
}
