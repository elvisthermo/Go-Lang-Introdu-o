package main

import "fmt"

func main() {
	numeros := []int{4, 32, 13, 34, 45}
	// Impressão dos elementos da matriz
	for i, numero := range numeros {
		fmt.Println("Index:", i, "Número:", numero)
	}

	// Criação de uma matriz bidimensional com 3 linhas e 4 colunas
    numerosMatriz := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

    // Impressão dos elementos da matriz
	
    for i, linha := range numerosMatriz {
        for j, numero := range linha {
            fmt.Println("Linha:", i, "Coluna:", j, "Número:", numero)
        }
    }

	fmt.Println("")
	fmt.Println("Valores Formatados:")
	for i := 0; i < 3; i++ {
        for j := 0; j < 4; j++ {
            fmt.Print(numerosMatriz[i][j] ,"\t")
        }
		fmt.Println();
    }

}
