package main

import (
	"time"
	"database/sql"
	"fmt"
	"os" // Importante para leer el sistema operativo

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Función auxiliar para leer variables de entorno con un valor por defecto
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func conectarDB() {
    var err error
    user := getEnv("DB_USER", "usuariodb")
    pass := getEnv("DB_PASS", "1234")
    host := getEnv("DB_HOST", "127.0.0.1")
    port := getEnv("DB_PORT", "3306")
    name := getEnv("DB_NAME", "proyecto_go")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)

    // Bucle de reintento: intentará 10 veces, una cada 2 segundos
    for i := 0; i < 10; i++ {
        db, err = sql.Open("mysql", dsn)
        if err == nil {
            err = db.Ping()
            if err == nil {
                fmt.Println("¡Conexión exitosa a la base de datos!")
                return // Sale de la función si tiene éxito
            }
        }
        fmt.Printf("Esperando a la base de datos en %s... (intento %d/10)\n", host, i+1)
        time.Sleep(2 * time.Second)
    }

    fmt.Println("Error crítico: No se pudo conectar a la DB tras varios intentos")
    os.Exit(1) // Detiene el programa si no hay conexión
}