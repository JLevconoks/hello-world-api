package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var startServing time.Time

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	delay := os.Getenv("STARTUP_DELAY")
	if delay == "" {
		delay = "0"
	}

	d, err := strconv.Atoi(delay)
	if err != nil {
		log.Println("'STARTUP_DELAY' should be numeric string in seconds, defaulting to '0'")
		d = 0
	}

	if d > 0 {
		log.Printf("Waiting for '%v' seconds startup delay", d)
	}

	startServing = time.Now().Add(time.Duration(d) * time.Second)

	router := http.NewServeMux()
	router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/ready", readyHandler)
	router.HandleFunc("/hello", helloHandler)

	log.Printf("Listening on :%v\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	if canServe() {
		statusCode = 200
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		statusCode = 400
	}
	w.WriteHeader(statusCode)
	log.Println("Hello request:", statusCode)

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	var statusCode int

	if canServe() {
		statusCode = 200
	} else {
		statusCode = 400
	}
	w.WriteHeader(statusCode)
	log.Println("Health request:", statusCode)
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	var statusCode int

	if canServe() {
		statusCode = 200
	} else {
		statusCode = 400
	}
	w.WriteHeader(statusCode)
	log.Println("Ready request:", statusCode)
}

func canServe() bool {
	return time.Now().After(startServing)
}
