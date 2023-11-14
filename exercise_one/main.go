/*
	Exercício 1

	Você deve avaliar 4 notas de um aluno (crie variáveis para o nome do aluno e para suas notas)
	Escreva uma função que retorne a avaliação desse aluno (o esqueleto já vou deixar escrito) abaixo
	Calcule a média do aluno e verifique:
  - Se a nota do aluno for menor que 5, retorne: NOME_DO_ALUNO: REPROVADO
	- Se a nota do aluno estiver entre 5 e 9, retorne: NOME_DO_ALUNO: APROVADO
	- Se a nota do aluno estiver acima de 9, retorne: NOME_DO_ALUNO: PARABÉNS

	Dica: Para pegar os inputs de nota do usuário, use a função fmt.Scan()
	- Para entender o fmt.Scan(): https://acervolima.com/funcao-fmt-scan-em-golang-com-exemplos/
	- Para entender como criar e escrever funções: https://go.dev/tour/basics/4
*/

// DISCLAIMER: FIZ DESSA FORMA PORQUE FICA FÁCIL DE TESTAR, QUALQUER DÚVIDA SÓ ME CHAMAREM

package main

import "fmt"

func main() {
	// Crie suas variáveis abaixo:

	var studentName string
	var grade1 float64
	var grade2 float64
	var grade3 float64
	var grade4 float64

	// Capture os inputs do usuário abaixo:

	fmt.Println("digite o nome do aluno")

	fmt.Scan(&studentName)

	fmt.Println("digite as notas do aluno")

	fmt.Scan(&grade1)
	fmt.Scan(&grade2)
	fmt.Scan(&grade3)
	fmt.Scan(&grade4)

	// Repasse suas variáveis para a função abaixo:

	result := reviewStudents(studentName, grade1, grade2, grade3, grade4)
	fmt.Println(result)

	// Printe no terminal o resultado esperado:
}

func reviewStudents(name string, grade1, grade2, grade3, grade4 float64) string {
	media := (grade1 + grade2 + grade3 + grade4) / 4

	if media < 5 {
		return fmt.Sprintf("%s: REPROVADO", name)
	} else if media >= 5 && media <= 9 {
		return fmt.Sprintf("%s: APROVADO", name)
	} else {
		return fmt.Sprintf("%s: PARABÉNS", name)
	}
}
