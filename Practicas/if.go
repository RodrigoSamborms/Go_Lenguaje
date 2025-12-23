package main
import "fmt"
func main(){
	// Sentencia if basica con else
	if 7%2 == 0 {
		fmt.Println("7 es par")
	}else {
		fmt.Println("7 es impar")
	}
	// sentencia if sin else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	// sentecia if con operadores lógicos
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("al menos 8 0 7 es par")
	}
	// Una sentencia puede preceder a condicionales, cualquier variable declarada en esta
	// esta disponible en todas las ramas siguientes del if, pero no fuera del if
	if num := 9; num < 0 {
		fmt.Println(num, "es negativo")
	}else if num < 10 {
		fmt.Println(num, "tiene un digito")
	}else {
		fmt.Println(num, "tiene multiples digitos")
	}
}