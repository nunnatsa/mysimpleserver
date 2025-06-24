package mux

import "net/http"

var m = http.NewServeMux()

func GetMux() *http.ServeMux {
	return m
}
