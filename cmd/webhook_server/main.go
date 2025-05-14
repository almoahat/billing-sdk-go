package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getPort()

	http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}

		fmt.Println("====== Webhook Received ======")
		fmt.Printf("Method: %s\n", r.Method)
		fmt.Printf("Headers: %v\n", r.Header)
		fmt.Println("Body:")
		fmt.Println(string(body))
		fmt.Println("==============================")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Printf("📡 Webhook server listening on http://localhost:%s/hook", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	return port
}
