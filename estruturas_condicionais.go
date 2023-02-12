package main

import "fmt"

func main(){
	x := 10
	// codição if
	if x > 0 {
		fmt.Println("x é positivo")
	 }

	y := 0
	// codição if else
	if y > 0 {
		fmt.Println("x é positivo")
	} else {
		fmt.Println("x é negativo ou zero")
	}

	// switch
	switch x := 5; x {
	case 1:
	   fmt.Println("x é 1")
	case 2, 3, 4:
	   fmt.Println("x é 2, 3 ou 4")
	default:
	   fmt.Println("x é maior que 4")
	}
}