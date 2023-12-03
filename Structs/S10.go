//Crie uma struct chamada Filme com os campos "título", "diretor", "ano" e "avaliações".
//O campo "avaliações" deve ser um slice de inteiros representando as notas que o filme
//recebeu dos espectadores. Escreva funções que permitam adicionar ou remover avaliações
//do filme, calcular a média das avaliações e imprimir as informações do filme e sua média de avaliações.

package main

import (
	"fmt"
)

type Filme struct {
	titulo      string
	diretor     string
	ano         int
	avaliacoes  []int
}

func (f *Filme) adicionarAvaliacao(avaliacao int) {
	f.avaliacoes = append(f.avaliacoes, avaliacao)
}

func (f *Filme) removerAvaliacao(index int) {
	if index >= 0 && index < len(f.avaliacoes) {
		f.avaliacoes = append(f.avaliacoes[:index], f.avaliacoes[index+1:]...)
	}
}

func (f *Filme) calcularMediaAvaliacoes() float64 {
	total := 0
	for _, avaliacao := range f.avaliacoes {
		total += avaliacao
	}
	if len(f.avaliacoes) > 0 {
		return float64(total) / float64(len(f.avaliacoes))
	}
	return 0.0
}

func (f *Filme) imprimirInfo() {
	fmt.Printf("Título: %s\n", f.titulo)
	fmt.Printf("Diretor: %s\n", f.diretor)
	fmt.Printf("Ano: %d\n", f.ano)
	fmt.Printf("Avaliações: %v\n", f.avaliacoes)
	fmt.Printf("Média de Avaliações: %.2f\n", f.calcularMediaAvaliacoes())
}

func main() {
	var filme Filme

	fmt.Print("Digite o título do filme: ")
	fmt.Scan(&filme.titulo)

	fmt.Print("Digite o diretor do filme: ")
	fmt.Scan(&filme.diretor)

	fmt.Print("Digite o ano do filme: ")
	fmt.Scan(&filme.ano)

	for {
		var avaliacao int
		fmt.Print("Digite uma avaliação (-1 para parar): ")
		fmt.Scan(&avaliacao)

		if avaliacao == -1 {
			break
		}

		filme.adicionarAvaliacao(avaliacao)
	}

	filme.imprimirInfo()
}
