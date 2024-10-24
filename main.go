package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "Status OK")
	if err != nil {
		log.Println("failed to write")
	}
}

func main() {
	http.HandleFunc("/status", HandlerStatus)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	data, err := json.Marshal("{}")
	if err != nil {
		log.Println("failed to marshal")
	}
	fmt.Println(data)
}
