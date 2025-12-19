package main
import "fmt" // import fmt package for formatted I/O
func main(){
	var a = "initial" // var declara una sola variable
	fmt.Println(a) // mostramos el valor de a
	var b, c int = 1, 2 // declaramos dos variables y las inicializamos
	fmt.Println(b,c) // mostramos el contendio de ambas variables
	var d = true // declaramos e inicializamos una variable booleana
	fmt.Println(d) // mostramos el contenido de d
	var e int // declaramos una variable tipo entera sin inicializar
	fmt.Println(e) // mostramos el contenido de e (valor por defecto 0)
	f := "apple" // la varialble ahora contiene un cadena y se infiere el tipo
	fmt.Println(f) //mostramos el contenido de la variable tipo cadena f
}