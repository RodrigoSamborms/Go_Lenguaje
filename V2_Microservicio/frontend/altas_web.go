package main

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func handlerAltasWeb(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        nombre := r.FormValue("nombre")
        
        // Creamos el objeto y lo convertimos a JSON
        datosUsuario := Usuario{Nombre: nombre}
        jsonDatos, _ := json.Marshal(datosUsuario)

        // Enviamos el JSON a la API
        //resp, err := http.Post("http://localhost:8081/api/usuarios/crear", 
        resp, err := http.Post("http://api-service:8081/api/usuarios/crear",
                               "application/json", 
                               bytes.NewBuffer(jsonDatos))

        if err != nil || resp.StatusCode != http.StatusCreated {
            http.Error(w, "Error al comunicarse con la API de guardado", 500)
            return
        }
    }
    // Después de guardar, regresamos al index
    http.Redirect(w, r, "/", http.StatusSeeOther)
}