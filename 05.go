package main

import "fmt"

// Crie um Array bidimensional de inteiros
//com 3 linhas e 2 colunas.
//Solicite ao usuário que informe os valores
//de cada elemento da matriz.
//Em seguida, imprima a matriz resultante.

func main() {
	var matriz [3][2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Digite o valor da linha %d e coluna %d: \n", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	for _, ints := range matriz {
		fmt.Println(ints)
	}
}
