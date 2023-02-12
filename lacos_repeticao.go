package main

import "fmt"

func main() {

	fmt.Println("Laço for")
	for i := 0; i < 5; i++ {
		fmt.Println("Index", i)
	}

	fmt.Println("Laço em array")
	nomes := []string{"John", "Jane", "Jim"}
    for i, nome := range nomes {
        fmt.Println("Index:", i, "Nome:", nome)
    }

	fmt.Println("Laço em While")
	// Em Go, o laço while é conseguido usando a sintaxe de laço 
	// for sem cláusulas de inicialização e incremento:
	i := 0
    for i < 5 {
        fmt.Println("Iteração", i)
        i++
    }

	fmt.Println("Laço for em Dicionários")
	idades := map[string]int{"John": 30, "Jane": 25, "Jim": 35}
    for nome, idade := range idades {
        fmt.Println("Nome:", nome, "Idade:", idade)
    }

}