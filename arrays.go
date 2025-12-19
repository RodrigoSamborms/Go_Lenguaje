package main
import "fmt"
func main() {
	//definición de un array de enteros con 5 elementos
	var a [5]int
	fmt.Println("emp: ", a)
	//ajuste de un valor especifico dentro del array usando su ínidice
	a[4] = 100
	fmt.Println("ajuste: ", a)
	fmt.Println("tomar: ", a[4])
	// podemos obtener la longitud del array usando la función len
	fmt.Println("lon: ", len(a))
	// la sitaxis siguiente declara y asigna valores al array en un solo paso
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// pudes dejar que el compilador asigne el número de elementos del array
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)
	// si especificas el ínidice con : los elementos entre serán zeros
	b =[...] int{100, 3: 400, 500}
	fmt.Println("idx:", b)
	// Se pueden crear array de varias dimensiones
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// se puede crear e inicializar un array de multiples dimensiones al mismo tiempo
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}