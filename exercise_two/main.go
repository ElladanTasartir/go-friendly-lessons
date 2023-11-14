/*
	Exercício 2

	Você deve coletar 5 números inteiros e identificar qual destes é o menor.
	Escreva uma função que retorne o menor número enviado nos parâmetros (o esqueleto já vou deixar escrito)

	Dica: Para pegar os inputs de nota do usuário, use a função fmt.Scan()
	- Para entender o fmt.Scan(): https://acervolima.com/funcao-fmt-scan-em-golang-com-exemplos/
	- Para entender como criar e escrever funções: https://go.dev/tour/basics/4
	- Para entender como criar e percorrer arrays: https://dev.to/linivecristine/arrays-em-golang-3cf1

*/

package main

import "fmt"

func main() {
	// Declare as variáveis aqui

	var num1 int
	var num2 int
	var num3 int
	var num4 int
	var num5 int

	// Colete os valores do usuário aqui

	fmt.Println("digite os números")

	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)
	fmt.Scan(&num4)
	fmt.Scan(&num5)

	// Chame a função lowestNumber aqui

	result := lowestNumber(num1, num2, num3, num4, num5)

	// Print o resultado da função aqui

	fmt.Println(result)

}

func lowestNumber(number1, number2, number3, number4, number5 int) int {
	// Verifique qual dos valores é o menor utilizando um laço de repetição e if

	numArray := [5]int{number1, number2, number3, number4, number5}
	fmt.Println(numArray, len(numArray))

	lowNum := numArray[0]

	for i := 0; i < len(numArray); i++ {
		if lowNum > numArray[i] {
			lowNum = numArray[i]
		}

	}

	// DICA: use um array de 5 posições para percorrer os valores enviados pelos parâmetros

	// Retorne o menor número encontrado

	return lowNum
}
