package main

import "fmt"

// Crie um programa que leia um array de inteiros
//e verifique se ele está ordenado em ordem crescente.

func main() {
	arr := [7]int{1, 2, 8, 4, 5, 6, 7}

	crescente := true

	for i := 0; i < len(arr)-1; i++ {
		atual := arr[i]
		prox := arr[i+1]
		if atual >= prox {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println("O array está ordenado")
	} else {
		fmt.Println("O array não está ordenado")
	}
}
