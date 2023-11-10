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
	var name string
	var grade1, grade2, grade3, grade4 float64

	// Capture os inputs do usuário abaixo:
	fmt.Println("Insira o nome da estudante:")
	fmt.Scan(&name)
	fmt.Println("Insira a primeira nota:")
	fmt.Scan(&grade1)
	fmt.Println("Insira a segunda nota:")
	fmt.Scan(&grade2)
	fmt.Println("Insira a terceira nota:")
	fmt.Scan(&grade3)
	fmt.Println("Insira a quarta nota:")
	fmt.Scan(&grade4)

	// Repasse suas variáveis para a função abaixo:
	averageGrade := reviewStudents(name, grade1, grade2, grade3, grade4)

	// Printe no terminal o resultado esperado:

	fmt.Print(averageGrade)
}

func reviewStudents(name string, grade1, grade2, grade3, grade4 float64) string {
	// Implemente o código responsável por calcular a média do aluno abaixo:
	average := (grade1 + grade2 + grade3 + grade4) / 4
	// Verifique a média do aluno e retorne a string esperada na descrição do exercício
	if average < 5 {
		return fmt.Sprintf("%v: REPROVADO", name)
	} else if average >= 5 && average < 9 {
		return fmt.Sprintf("%v: APROVADO", name)
	} else {
		return fmt.Sprintf("%v: PARABÉNS", name)
	}
}
