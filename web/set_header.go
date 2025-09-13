package web

import "net/http"

func JsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func htmlHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "Text/Html; charset=utf-8")
}
