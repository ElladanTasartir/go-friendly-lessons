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
	var number1, number2, number3, number4, number5 int

	// Colete os valores do usuário aqui
	fmt.Println("Insira o primeiro número inteiro:")
	fmt.Scan(&number1)
	fmt.Println("Insira o primeiro número inteiro:")
	fmt.Scan(&number2)
	fmt.Println("Insira o primeiro número inteiro:")
	fmt.Scan(&number3)
	fmt.Println("Insira o primeiro número inteiro:")
	fmt.Scan(&number4)
	fmt.Println("Insira o primeiro número inteiro:")
	fmt.Scan(&number5)

	// Chame a função lowestNumber aqui
	lowestNumberResult := lowestNumber(number1, number2, number3, number4, number5)

	// Print o resultado da função aqui
	fmt.Printf("%v is the minimum number.", lowestNumberResult)
}

func lowestNumber(number1, number2, number3, number4, number5 int) int {
	// Verifique qual dos valores é o menor utilizando um laço de repetição e if
	// DICA: use um array de 5 posições para percorrer os valores enviados pelos parâmetros
	numberArray := [5]int{number1, number2, number3, number4, number5}
	minimumNumber := numberArray[0]

	for i := 0; i < len(numberArray); i++ {
		if numberArray[i] < minimumNumber {
			minimumNumber = numberArray[i]
		}
	}

	// Retorne o menor número encontrado
	return minimumNumber
}
