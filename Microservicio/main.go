package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed index.html styles.css
var archivosWeb embed.FS

func main() {
    conectarDB()

    // 1. Página principal
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        contenido, _ := archivosWeb.ReadFile("index.html")
        w.Header().Set("Content-Type", "text/html")
        w.Write(contenido)
    })

    // 2. RECUPERAR EL CSS (Esto es lo que faltaba)
    http.HandleFunc("/static/styles.css", func(w http.ResponseWriter, r *http.Request) {
        contenido, _ := archivosWeb.ReadFile("styles.css")
        w.Header().Set("Content-Type", "text/css")
        w.Write(contenido)
    })

    // 3. Handlers modulares
    http.HandleFunc("/agregar", handlerAltas)
    http.HandleFunc("/listar", handlerListas)
	// En main.go, dentro de func main()
	http.HandleFunc("/borrar", handlerBorrar)
	http.HandleFunc("/editar", handlerEditar)

    fmt.Println("Servidor en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}