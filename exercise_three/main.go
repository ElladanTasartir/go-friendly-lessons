/*
	Exercício 3

	Você deve analisar uma lista de valores inteiros. Identifique se algum dos valores se repete nesta lista:
	- Se sim retorne verdade
	- Se falso retorne falso

	Dica:
	- Para entender como criar e percorrer arrays: https://dev.to/linivecristine/arrays-em-golang-3cf1
*/

package main

import (
	"fmt"
)

func main() {
	// Crie aqui a sua lista

	listNumber := 0
	arr := []int{}

	// Pergunte para o usuário quantos valores ele quer inserir na lista

	fmt.Println("Quais valores você deseja inserir?")
	fmt.Scan(&listNumber)

	// Insira cada valor que o usuário vai digitar em uma lista

	fmt.Println("insira os valores aqui")
	fmt.Print(&arr)

	for i := 0; i < listNumber; i++ {
		arrNumber := 0
		fmt.Scan(&arrNumber)
		arr = append(arr, arrNumber)

	}

	// Repasse os valores para a função abaixo
	// Print o valor retornado da função

	fmt.Print(hasRepeatingValue(arr))

}

func hasRepeatingValue(numbers []int) bool {
	// Percorra o array e procure repetições entre valores
	for i, number := range numbers {
		for j, compareNumber := range numbers {
			if i != j && number == compareNumber {
				return true
			}
		}
	}

	// Retorne falso caso não ocorra

	return false
}
