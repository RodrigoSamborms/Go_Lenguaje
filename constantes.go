package main
import (
	"fmt"
	"math"
)
const s string = "constant" // declaramos una constante de tipo cadena
func main(){
	fmt.Println(s) //mostramos la constante s que fue creada "fuera de una funcion"
	const n = 500000000 // declaramos una constante entera muy grande
	const d = 3e20 / n // declaramos una constante flotante muy pequeña
	fmt.Println(d) // mostramos el valor de d
	fmt.Println(int64(d)) // d se convierte en entero de 64 bits y se muestra
	fmt.Println(math.Sin(n)) // n se vuelve un enetero de 64 bits como espera la funcion Sin
}