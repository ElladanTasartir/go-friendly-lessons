/*
	Exercício 3

	Você deve analisar uma lista de valores inteiros. Identifique se algum dos valores se repete nesta lista:
	- Se sim retorne verdade
	- Se falso retorne falso

	Dica:
	- Para entender como criar e percorrer arrays: https://dev.to/linivecristine/arrays-em-golang-3cf1
*/

package main

import "fmt"

func main() {
	// Crie aqui a sua lista
	loopCount := 0
	intArray := []int{}

	// Pergunte para o usuário quantos valores ele quer inserir na lista
	fmt.Println("Insira a quantidade de valores da lista:")
	fmt.Scan(&loopCount)

	// Insira cada valor que o usuário vai digitar em uma lista
	fmt.Println("Insira respectivamente os valores individuais da lista:")

	for i := 0; i < loopCount; i++ {
		arrayValue := 0
		fmt.Scan(&arrayValue)
		intArray = append(intArray, arrayValue)
	}

	// Repasse os valores para a função abaixo
	// Print o valor retornado da função
	fmt.Print(hasRepeatingValue(intArray))
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
