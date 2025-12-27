package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura idéntica a la del Frontend para que hablen el mismo idioma
type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func main() {
	conectarDB() // Función que ya tienes en db.go

	// Endpoint que el Frontend está intentando llamar
	http.HandleFunc("/api/usuarios", func(w http.ResponseWriter, r *http.Request) {
		usuarios := obtenerUsuariosDeDB()
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(usuarios)
	})

	// Agrega esto en el main() de main_api.go
	http.HandleFunc("/api/usuarios/crear", crearUsuarioAPI)

	// Agrega esto en el main() de main_api.go
	http.HandleFunc("/api/usuarios/borrar", borrarUsuarioAPI)

    // En el main() del backend
    http.HandleFunc("/api/usuarios/editar", editarUsuarioAPI)

	fmt.Println("Backend API corriendo en http://localhost:8081")
    //fmt.Println("🚀 Frontend iniciado en http://localhost:8080")
	http.ListenAndServe(":8081", nil)
}

// Función para procesar el listado desde la DB
func obtenerUsuariosDeDB() []Usuario {
	var lista []Usuario
	rows, err := db.Query("SELECT id, nombre FROM usuarios WHERE activo = 1")
	if err != nil {
		fmt.Println("Error en query:", err)
		return lista
	}
	defer rows.Close()

	for rows.Next() {
		var u Usuario
		if err := rows.Scan(&u.ID, &u.Nombre); err == nil {
			lista = append(lista, u)
		}
	}
	return lista
}

// Función para procesar el guardado
func crearUsuarioAPI(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", 405)
        return
    }

    var nuevo Usuario
    // Decodificamos el JSON que viene del Frontend
    err := json.NewDecoder(r.Body).Decode(&nuevo)
    if err != nil || nuevo.Nombre == "" {
        http.Error(w, "Datos inválidos", 400)
        return
    }

    // Insertamos en la DB (usando tu lógica de db.go)
    _, err = db.Exec("INSERT INTO usuarios (nombre) VALUES (?)", nuevo.Nombre)
    if err != nil {
        http.Error(w, "Error al insertar en DB", 500)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, `{"mensaje": "Usuario creado con éxito"}`)
}

// Función para procesar el borrado lógico
func borrarUsuarioAPI(w http.ResponseWriter, r *http.Request) {
    // Obtenemos el ID de la URL (ej: /api/usuarios/borrar?id=5)
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID requerido", 400)
        return
    }

    // Ejecutamos el borrado lógico (UPDATE activo = 0)
    _, err := db.Exec("UPDATE usuarios SET activo = FALSE WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Error al actualizar en DB", 500)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, `{"mensaje": "Usuario desactivado"}`)
}

// Función para procesar la edición
func editarUsuarioAPI(w http.ResponseWriter, r *http.Request) {
    // 1. Extraer ID de la URL (?id=...)
    id := r.URL.Query().Get("id")
    
    // 2. Leer el JSON del cuerpo
    var datos struct {
        Nombre string `json:"nombre"`
    }
    
    err := json.NewDecoder(r.Body).Decode(&datos)
    if err != nil {
        fmt.Println("Error decodificando JSON:", err)
        http.Error(w, "JSON inválido", 400)
        return
    }

    // 3. Ejecutar en DB
    if id != "" && datos.Nombre != "" {
        _, err := db.Exec("UPDATE usuarios SET nombre = ? WHERE id = ?", datos.Nombre, id)
        if err != nil {
            fmt.Println("Error en SQL:", err)
            http.Error(w, "Error de base de datos", 500)
            return
        }
        fmt.Printf("Usuario %s actualizado a %s\n", id, datos.Nombre)
        w.WriteHeader(http.StatusOK)
    }
}