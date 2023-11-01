package main

import (
	"fmt"
	"net/http"
)

func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Status OK")
}
