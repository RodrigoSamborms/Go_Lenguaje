package main
import (
	"fmt"
	"maps"
)

func main(){
	//declaramos un map vacio usando la funcion make
	m := make(map[string]int) //sintaxys es make(map[tipoClave]tipoValor)
	//asignasr valores al mapa se realiza con la sintaxys de pares clave-valor
	//name[clave] = valor
	m["k1"]= 7
	m["k2"]= 13
	//imprimir un mapa mostrara los pares clave-valor
	fmt.Println("map:", m)
	//conseguimos un valor del mapa usando su clave en lugar del indice como en los arrays
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	//si la clave no existe en el mapa se devuelve el valor cero
	v3 := m["k3"]
	fmt.Println("v3:", v3)
	//la funcion len aplicada sobre un mapa regresa el número de valores clave-valor en el mapa
	fmt.Println("len:", len(m))
	//la función delete elemina un par clave-valor del mapa
	delete(m, "k2")
	fmt.Println("map:", m)
	//para eliminar todos los valores dentro del mapa se utiliza la funcion clear
	clear(m)
	fmt.Println("map:", m)
	//retornar el segundo valor es opcional, esto puede ser util para desambiguar
	//entre claves perdidas y claves con valores cero o "" vacias.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	//también es posible inicializar un mapa con valores al declararlo
	n := map[string]int{"foo" : 1, "bar": 2}
	fmt.Println("map:", n)
	//revise el paquete maps para más funciones útiles
	n2 := map[string]int{"foo" : 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}