package main

import "fmt"

// Crie um Slice de inteiros e solicite ao usuário
//que informe o tamanho do Slice e os valores
//dos elementos. Em seguida, imprima o Slice
//e a soma dos valores armazenados.

func main() {
	var tamanho, valor int

	fmt.Print("Digite o tamanho da Slice: ")
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite um valor: ")
		fmt.Scan(&valor)
		slice[i] = valor
	}

	var soma int

	for i := 0; i < len(slice); i++ {
		x := slice[i]
		soma += x
	}

	fmt.Println("Meu slice é: ", slice)
	fmt.Println("A soma é: ", soma)
}
