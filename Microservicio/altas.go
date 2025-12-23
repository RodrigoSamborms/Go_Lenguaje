package main

import "net/http"

func handlerAltas(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nombre := r.FormValue("nombre")
		if nombre != "" {
			_, err := db.Exec("INSERT INTO usuarios (nombre) VALUES (?)", nombre)
			if err != nil {
				http.Error(w, "Error al insertar", 500)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}