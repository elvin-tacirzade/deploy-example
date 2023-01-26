package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", testHandler)
	lstErr := http.ListenAndServe(":8008", nil)
	log.Fatalf("failed to listen: %v", lstErr)
}

func testHandler(w http.ResponseWriter, _ *http.Request) {
	_, writeErr := w.Write([]byte("Hello World"))
	if writeErr != nil {
		log.Fatalf("failed to write: %v", writeErr)
	}
}
