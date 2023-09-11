package main

import "fmt"

// Crie um Slice de inteiros com o tamanho 5.
//Em seguida, solicite ao usuário que informe
//um número inteiro.
//Adicione esse número ao Slice apenas se ele não
//estiver presente.
//Imprima o Slice resultante.

func main() {
	slice := []int{1, 2, 3, 4, 5}

	var x int

	fmt.Print("Digite um valor: ")
	fmt.Scan(&x)

	var achou bool

	for _, valor := range slice {
		if x == valor {
			achou = true
			break
		}
	}

	if !achou {
		slice = append(slice, x)
	}
}
