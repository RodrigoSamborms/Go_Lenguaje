package main
import "fmt"
func main(){
	i:= 1 //el for mas sencillo de implementar excepto el infinito
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	} 
	for j := 0; j < 3; j++ { //el for clasico inicio/condición/incremento
		fmt.Println(j)
	}
	for i := range 3 { // el basico "haz esto N veces"
		fmt.Println("range", i)
	}
	for {		// el for se repetira hasta encontrar un break o return
		fmt.Println("loop")
		break
	}
	for n := range 6 { //el lazo continua mientras se ejecute la instrucción continue
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}