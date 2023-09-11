package main

import "fmt"

// Crie um Slice de strings com tamanho 8 e
//solicite ao usuário que informe um valor.
//Remova todas as ocorrências desse valor no Slice
//e imprima o resultado.

func main() {
	var slice = []string{"a", "b", "c", "10", "e", "f", "g", "h"}
	var x string
	fmt.Print("Digite um valor: ")
	fmt.Scan(&x)

	for i := 0; i < len(slice); i++ {
		if slice[i] == x {
			fmt.Println(i, x, slice[i])
			slice = append(slice[:i], slice[i+1:]...)
		}
	}

	fmt.Println(slice)
}
