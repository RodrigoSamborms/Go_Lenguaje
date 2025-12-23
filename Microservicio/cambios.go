package main

import (
	"net/http"
)

// Borrado Lógico: Solo cambia el estado a 'false'
func handlerBorrar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		_, err := db.Exec("UPDATE usuarios SET activo = FALSE WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Error al borrar", 500)
			return
		}
	}
	http.Redirect(w, r, "/listar", http.StatusSeeOther)
}

// Edición: En este ejemplo simple, cambiaremos el nombre por uno predefinido o nuevo
func handlerEditar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	nuevoNombre := r.FormValue("nuevoNombre")

	if id != "" && nuevoNombre != "" {
		_, err := db.Exec("UPDATE usuarios SET nombre = ? WHERE id = ?", nuevoNombre, id)
		if err != nil {
			http.Error(w, "Error al editar", 500)
			return
		}
	}
	http.Redirect(w, r, "/listar", http.StatusSeeOther)
}