package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func handlerListasWeb(w http.ResponseWriter, r *http.Request) {
	// 1. Configuramos un cliente con "timeout" por si la API tarda mucho
	cliente := http.Client{Timeout: 5 * time.Second}
	
	resp, err := cliente.Get(apiURL)
	if err != nil {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h2>Error: No se pudo conectar con la API</h2><p>%v</p><a href='/'>Volver</a>", err)
		return
	}
	defer resp.Body.Close()

	// 2. Intentamos leer el JSON
	var usuarios []Usuario
	err = json.NewDecoder(resp.Body).Decode(&usuarios)
	if err != nil {
		http.Error(w, "Error al procesar los datos de la API", 500)
		return
	}

	// 3. Renderizamos el HTML
	renderizarTabla(w, usuarios)
}

func renderizarTabla(w http.ResponseWriter, usuarios []Usuario) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<html><head><link rel="stylesheet" href="/static/styles.css"></head><body><div class="seccion">`)
	fmt.Fprint(w, "<h1>Usuarios Activos (Desde Microservicio API)</h1><ul>")
	
	for _, u := range usuarios {
		fmt.Fprintf(w, `
			<li>
				<span class="nombre-usuario">[%d] %s</span>
				<div class="acciones">
					<a href="/borrar?id=%d" class="btn btn-borrar">Borrar</a>
					
					<form action="/editar?id=%d" method="POST" style="display:inline;">
						<input type="text" name="nuevoNombre" placeholder="Nuevo nombre" required>
						<button type="submit" class="btn btn-editar">Editar</button>
					</form>
				</div>
			</li>`, u.ID, u.Nombre, u.ID, u.ID)
	}
	
	fmt.Fprint(w, "</ul><br><hr><a href='/' class='btn'>← Regresar</a></div></body></html>")
}