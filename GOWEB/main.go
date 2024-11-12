// main.go
package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>¡Hola Mundo desde Go!</h1>")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Servidor en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
