package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed index.html styles.css
var archivosWeb embed.FS

// Esta será la dirección interna en tu red de Docker/Kubernetes después
// Por ahora, para pruebas locales, usaremos localhost
//const apiURL = "http://localhost:8081/api/usuarios"  //para ejecutar localmente
const apiURL = "http://api-service:8081/api/usuarios" //ejecución en Docker

func main() {
    // Servir la página de inicio
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        contenido, _ := archivosWeb.ReadFile("index.html")
        w.Header().Set("Content-Type", "text/html")
        w.Write(contenido)
    })

    // Servir el CSS
    http.HandleFunc("/static/styles.css", func(w http.ResponseWriter, r *http.Request) {
        contenido, _ := archivosWeb.ReadFile("styles.css")
        w.Header().Set("Content-Type", "text/css")
        w.Write(contenido)
    })

    // Handlers que "puentean" a la API
    http.HandleFunc("/listar", handlerListasWeb)
    // http.HandleFunc("/agregar", handlerAltasWeb) // Lo habilitaremos en el siguiente paso
	http.HandleFunc("/agregar", handlerAltasWeb) // Habilitado para el siguiente paso
	// http.HandleFunc("/borrar", handlerBorrarWeb) // Lo habilitaremos en el siguiente paso
	http.HandleFunc("/borrar", handlerBorrarWeb)  // Habilitado para el siguiente paso
    // http.HandleFunc("/editar", handlerEditarWeb) // Lo habilitaremos en el siguiente paso
    http.HandleFunc("/editar", handlerEditarWeb) // Habilitado para el siguiente paso

    fmt.Println("🚀 Frontend iniciado en http://localhost:8080")
    //fmt.Println("Backend API corriendo en http://localhost:8081")
    
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error al iniciar el servidor: %v\n", err)
    }
}