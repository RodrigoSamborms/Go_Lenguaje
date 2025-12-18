# Ejecución de un programa en Go
Asegure de estar en WSL Debian, copie el siguiente código en el archivo hello-world.go
```
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```
Guarde el archivo, para ejecutarlo, o crear el ejecutable use la siguientes líneas según el caso:

```
~$go run hello-worl.go #para ejecutar el archivo

~$go build hello-world.go
~$ls
#verifique la creación del archivo y ejecute con la siguiente linea
~$./hello-world

```

 ![Primer Programa](./imagenes/Go_001.jpg)