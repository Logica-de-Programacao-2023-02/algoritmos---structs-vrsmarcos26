//Crie uma struct chamada Viagem com os campos "origem", "destino", "data" e "preço".
//Escreva uma função que receba um slice de Viagens como parâmetro e retorne a viagem mais cara.

package main

import "fmt"

type Viagem struct {
	origem  string
	destino string
	data    string
	preco   int
}

var preco1 int
var preco2 int
var preco3 int
var preco4 int
var maiscaro string
var top1 int

var lugareseprecos = []string{"A", "B", "C", "D"}

func viagemmaiscara() {
	if preco1 > preco2 && preco1 > preco3 && preco1 > preco4 {
		maiscaro = "A"
	}
	if preco2 > preco1 && preco2 > preco3 && preco2 > preco4 {
		maiscaro = "B"
	}
	if preco4 > preco1 && preco4 > preco2 && preco4 > preco3 {
		maiscaro = "D"
	}
	if preco3 > preco1 && preco3 > preco2 && preco3 > preco4 {
		maiscaro = "C"
	}
}

func preçomaiscaro() {
	if preco1 > preco2 && preco1 > preco3 && preco1 > preco4 {
		top1 = preco1
	}
	if preco2 > preco1 && preco2 > preco3 && preco2 > preco4 {
		top1 = preco2
	}
	if preco4 > preco1 && preco4 > preco2 && preco4 > preco3 {
		top1 = preco4
	}
	if preco3 > preco1 && preco3 > preco2 && preco3 > preco4 {
		top1 = preco3
	}
}

func main() {
	fmt.Print("Preço do destino A: ")
	fmt.Scan(&preco1)
	fmt.Print("Preço do destino B: ")
	fmt.Scan(&preco2)
	fmt.Print("Preço do destino C: ")
	fmt.Scan(&preco3)
	fmt.Print("Preço do destino D: ")
	fmt.Scan(&preco4)
	fmt.Print("ENTÃO a viagem mais barata é: ")
	viagemmaiscara()
	preçomaiscaro()
	fmt.Print(Viagem{origem: "Brasil", destino: maiscaro, data: "hoje", preco: top1})
}
