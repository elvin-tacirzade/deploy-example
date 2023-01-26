package main

import (
	"log"
	"net/http"
)

const PORT = ":8008"

func main() {
	http.HandleFunc("/test", testHandler)
	log.Printf("http server started on: %v", PORT)
	lstErr := http.ListenAndServe(PORT, nil)
	log.Fatalf("failed to listen: %v", lstErr)
}

func testHandler(w http.ResponseWriter, _ *http.Request) {
	_, writeErr := w.Write([]byte("Hello World"))
	if writeErr != nil {
		log.Fatalf("failed to write: %v", writeErr)
	}
}
