package main
import (
	"fmt"
	"slices"
)
func main(){
	//definición de un slice de tamaño cero
	var s []string
	fmt.Println("unint:", s ,s == nil, len(s) == 0)
	//para crear un slice con longitud cero usa make
	s = make([]string, 3)
	fmt.Println("emp:", s, "len: ", len(s), "cap: ", cap(s))
	//darle valores al slice es igual que con un array
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get.", s[2])
	//len regresa la lingitud del slice
	fmt.Println("len:", len(s))
	//los slice a diferencia de los arreglos puede aceptar nuevos elementos
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("adp:", s)
	//podemos copiar un slice a otro siempre que tengan la misma longitud
	c := make([]string, len(s)) //construimos un slice con la misma longitud de s
	copy(c,s)
	fmt.Println("cpy:", c)
	//los slices pueden ser rebanados (sliced)
	l := s[2:5]
	fmt.Println("sl1:", l)
	//este slice toma los valores de s dejando el 5 fuera
	l = s[:5]
	fmt.Println("sl2:", l)
	//este slice toma los valores desde el incio hasta el 2
	l = s[2:]
	fmt.Println("sl3:", l)
	//también podemos declarar e inicializar el slice en una sola línea
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	//existen varias funciones utiles en el paquete slices
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2){
		fmt.Println("t == t2")
	}
	//slices pueden se compuestos en estructuras multidimensionales, excepto que 
	//a diferencia de los arrays, estos pueden ser redimensionados
	twoD := make([][]int, 3)	//slice de 3 elementos y dos slices
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}