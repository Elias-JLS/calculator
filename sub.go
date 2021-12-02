package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func sub(a, b float64) float64 {
	return a - b
}

func subHandler(
	w http.ResponseWriter,
	r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	num1, err := strconv.ParseFloat(partes[2], 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	num2, err := strconv.ParseFloat(partes[3], 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	result := sub(num1, num2)
	fmt.Fprintf(w, "%.2f", result)
}