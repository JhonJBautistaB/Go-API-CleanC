package handlers

import (
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello mi llave, implemento arquitectura limmpia"))
}

func PruebaWeb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hiciste otro endpoint!!! felicitaciones"))
}
