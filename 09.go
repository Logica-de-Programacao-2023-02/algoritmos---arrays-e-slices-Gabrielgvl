package main

import "fmt"

// Crie um Array de floats com 6 elementos.
//Solicite ao usuário que informe um número
//que será adicionado em todas as posições do Array.
//Imprima o Array resultante.

func main() {
	arr := [6]float64{1.4, 1.6, 1.8, 2.2, 2.4, 2.6}

	var x float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&x)

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + x
	}

	fmt.Println(arr)
}
