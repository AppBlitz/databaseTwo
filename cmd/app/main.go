package main

import (
	"fmt"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el parámetro "msg" de la URL: ?msg=hola
	msg := r.URL.Query().Get("msg")

	if msg == "" {
		http.Error(w, "Falta el parámetro 'msg'", http.StatusBadRequest)
		return
	}

	// Devolver el mismo mensaje
	fmt.Fprintf(w, "Recibido: %s", msg)
}

func main() {
	http.HandleFunc("/echo", echoHandler)

	fmt.Println("Servidor escuchando en http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
