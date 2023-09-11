package main

import "fmt"

// Faça um programa que leia dois arrays de
//inteiros de tamanho n e gere um terceiro
//array que seja a soma dos dois primeiros.

func main() {
	var n int

	fmt.Print("Digite o tamanho das slices: ")
	fmt.Scan(&n)

	slice1 := make([]int, n)
	slice2 := make([]int, n)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Digite o valor %d da slice 1: ", i)
		fmt.Scan(&slice1[i])

		fmt.Printf("Digite o valor %d da slice 2: ", i)
		fmt.Scan(&slice2[i])
	}

	for i := 0; i < n; i++ {
		result[i] = slice1[i] + slice2[i]
	}

	fmt.Println(result)
}
