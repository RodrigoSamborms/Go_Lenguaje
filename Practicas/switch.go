package main
import (
	"fmt"
	"time"
)
func main() {
	// Sentencia switch simple
	i := 2
	fmt.Print("Escribe ", i, " como " )
	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	}
	// Senecias multiples, separadas por comas
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("Es fin de semana")
	default:
		fmt.Println("Es un día entre semana")
	}
	// El switch sin expresión sirve para expresar logica de if-else las opciones pueden ser
	// no constantes
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println(" es antes del mediodía")
	default:
		fmt.Println("es después del mediodía")
	}
	// también se pueden comparar tipos de datos en variables
	// esto puede usarse para descubir el tipo dinámico de una interfaz
	queSoyYo := func(i interface{}) {
		switch t := i.(type){
		case bool:
			fmt.Println("Soy un booleano")
		case int:
			fmt.Println("Soy un entero")
		default:
			fmt.Printf("No se que tipo soy &T\n", t)
		}
	}
	// ejemplos de la función para para diferentes tipos de datos
	queSoyYo(true)
	queSoyYo(1)
	queSoyYo("hey")
}