package main

import (
	"bytes"   // <--- Faltaba esta
	"encoding/json" // <--- Faltaba esta
	"fmt"
	"net/http"
)

//handler para borrar usuarios logicamente
func handlerBorrarWeb(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id != "" {
        // Construimos la URL de la API con el ID
        //urlAPI := fmt.Sprintf("http://localhost:8081/api/usuarios/borrar?id=%s", id)
        urlAPI := fmt.Sprintf("http://api-service:8081/api/usuarios/borrar?id=%s", id)
        
        // Usamos el método DELETE o un simple GET/POST según prefieras, 
        // para simplificar usaremos un NewRequest
        req, _ := http.NewRequest(http.MethodPut, urlAPI, nil) 
        cliente := &http.Client{}
        resp, err := cliente.Do(req)

        if err != nil || resp.StatusCode != http.StatusOK {
            http.Error(w, "No se pudo borrar el usuario en la API", 502)
            return
        }
    }
    // Volvemos a la lista para ver los cambios
    http.Redirect(w, r, "/listar", http.StatusSeeOther)
}

//handler para modificar usuarios
func handlerEditarWeb(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    nuevoNombre := r.FormValue("nuevoNombre")

    if id != "" && nuevoNombre != "" {
        // 1. Crear el JSON
        cuerpoJSON, _ := json.Marshal(map[string]string{"nombre": nuevoNombre})
        
        // 2. Preparar la petición a la API (Puerto 8081)
        //url := fmt.Sprintf("http://localhost:8081/api/usuarios/editar?id=%s", id)
        url := fmt.Sprintf("http://api-service:8081/api/usuarios/editar?id=%s", id)
        req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(cuerpoJSON))
        
        // ¡IMPORTANTE!: Avisar a la API que mandamos JSON
        req.Header.Set("Content-Type", "application/json")

        cliente := &http.Client{}
        resp, err := cliente.Do(req)
        
        // Debug temporal: Si falla, muestra el error en lugar de redireccionar
        if err != nil || resp.StatusCode != http.StatusOK {
            fmt.Printf("Error llamando a la API: %v\n", err)
            http.Error(w, "La API de backend no procesó el cambio", 502)
            return
        }
    }
    http.Redirect(w, r, "/listar", http.StatusSeeOther)
}