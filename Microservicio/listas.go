package main

import (
	"fmt"
	"net/http"
)

func handlerListas(w http.ResponseWriter, r *http.Request) {
	// Filtramos solo los activos
	rows, err := db.Query("SELECT id, nombre FROM usuarios WHERE activo = TRUE")
	if err != nil {
		http.Error(w, "Error al consultar", 500)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<html><head><link rel="stylesheet" href="/static/styles.css"></head><body><div class="seccion">`)
	fmt.Fprint(w, "<h1>Usuarios Activos</h1><ul>")
	
	for rows.Next() {
		var id int
		var nombre string
		rows.Scan(&id, &nombre)
		
		// Añadimos botones de Editar y Borrar para cada fila
		// Dentro del bucle for rows.Next() en listas.go
		fmt.Fprintf(w, `
			<li>
				<span class="nombre-usuario">%s</span> 
				<div class="acciones">
					<a href="/borrar?id=%d" class="btn btn-borrar">Borrar</a>

					<form action="/editar?id=%d" method="POST" style="display:inline;">
						<input type="text" name="nuevoNombre" placeholder="Nuevo nombre" required>
						<button type="submit" class="btn btn-editar">Editar</button>
					</form>
				</div>
			</li>`, nombre, id, id)
	}
	
	fmt.Fprint(w, `
	<html>
    <head>
        <meta charset="UTF-8"> <title>Lista de Usuarios</title>
        <link rel="stylesheet" href="/static/styles.css">
    </head>
    <body>
            </ul>
            <br>
            <hr>
            <a href="/" class="btn btn-regresar">← Regresar al formulario</a>
        </div>
    </body>
    </html>`)
}