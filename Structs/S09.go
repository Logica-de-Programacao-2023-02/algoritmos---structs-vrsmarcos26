//Crie uma struct chamada Aluno com os campos "nome", "idade" e "notas".
//O campo "notas" deve ser um slice de float64, representando as notas
//que o aluno tirou em uma determinada disciplina. Escreva funções que
//permitam adicionar ou remover notas do aluno, calcular a média das notas e imprimir o nome, idade e média do aluno.

package main

import (
	"fmt"
)

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func (a *Aluno) adicionarNota(nota float64) {
	a.notas = append(a.notas, nota)
}

func (a *Aluno) calcularMedia() float64 {
	total := 0.0
	for _, nota := range a.notas {
		total += nota
	}
	if len(a.notas) > 0 {
		return total / float64(len(a.notas))
	}
	return 0.0
}

func (a *Aluno) imprimirInfo() {
	fmt.Printf("Nome: %s\n", a.nome)
	fmt.Printf("Idade: %d\n", a.idade)
	fmt.Printf("Notas: %v\n", a.notas)
	fmt.Printf("Média: %.2f\n", a.calcularMedia())
}

func main() {
	var aluno Aluno

	fmt.Print("Diga seu nome: ")
	fmt.Scan(&aluno.nome)

	fmt.Print("Agora diga sua idade: ")
	fmt.Scan(&aluno.idade)

	for {
		var nota float64
		fmt.Print("Me diga sua nota (-1 para parar): ")
		fmt.Scan(&nota)

		if nota == -1 {
			break
		}

		aluno.adicionarNota(nota)
	}

	aluno.imprimirInfo()
}
