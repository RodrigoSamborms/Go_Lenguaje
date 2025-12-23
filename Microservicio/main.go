package main

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Driver de MariaDB
)

//go:embed index.html styles.css
var archivosWeb embed.FS

// Variable global para la base de datos (Pool de conexiones)
var db *sql.DB

func main() {
	var err error
	// Configuración de conexión: "usuario:password@tcp(127.0.0.1:3306)/nombre_db"
	// Ajusta estos datos según tu configuración de MariaDB
	dsn := "usuariodb:1234@tcp(127.0.0.1:3306)/proyecto_go"
	
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error al configurar la DB: %v\n", err)
		return
	}
	
	// Verificar si la conexión es real
	err = db.Ping()
	if err != nil {
		fmt.Printf("No se pudo conectar a MariaDB: %v\n", err)
		return
	}
	fmt.Println("Conexión exitosa a MariaDB")

	// --- HANDLERS ---

	// 1. Inicio (Servir HTML)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		contenido, _ := archivosWeb.ReadFile("index.html")
		w.Header().Set("Content-Type", "text/html")
		w.Write(contenido)
	})

	// 2. Estilos (Servir CSS)
	http.HandleFunc("/static/styles.css", func(w http.ResponseWriter, r *http.Request) {
		contenido, _ := archivosWeb.ReadFile("styles.css")
		w.Header().Set("Content-Type", "text/css")
		w.Write(contenido)
	})

	// 3. AGREGAR (C de CRUD)
	http.HandleFunc("/agregar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			nombre := r.FormValue("nombre")
			
			// Insertar en MariaDB
			_, err := db.Exec("INSERT INTO usuarios (nombre) VALUES (?)", nombre)
			if err != nil {
				http.Error(w, "Error al guardar en la DB", 500)
				return
			}
			
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}