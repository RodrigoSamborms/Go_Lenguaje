package main

import (
	"embed"
	"fmt"
	"net/http"
)

// Usamos un patrón (wildcard) para embeber todos los archivos .html y .css
//go:embed index.html styles.css
var archivosWeb embed.FS

func main() {
	// Ruta para el HTML principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		contenido, _ := archivosWeb.ReadFile("index.html")
		w.Header().Set("Content-Type", "text/html")
		w.Write(contenido)
	})

	// Ruta para el CSS
	// Cuando el HTML pida "/static/styles.css", Go lo servirá desde el binario
	http.HandleFunc("/static/styles.css", func(w http.ResponseWriter, r *http.Request) {
		contenido, _ := archivosWeb.ReadFile("styles.css")
		w.Header().Set("Content-Type", "text/css")
		w.Write(contenido)
	})

	// ... (tus otros handlers de /agregar y /listar permanecen igual) ...

	fmt.Println("Servidor con CSS separado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}