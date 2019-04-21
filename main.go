package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := http.NewServeMux()
	router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/hello", helloHandler)

	log.Printf("Listening on :%v\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
