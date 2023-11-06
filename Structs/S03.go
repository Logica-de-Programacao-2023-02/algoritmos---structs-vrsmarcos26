//Crie uma struct chamada Triângulo com os campos "base" e "altura".
//Escreva uma função que receba um Triângulo como parâmetro e calcule a área do triângulo (área = base * altura / 2).

package main

import "fmt"

type Triangulo struct {
	base   int
	altura int
}

func main() {
	p := Triangulo{base: 2, altura: 4}
	a := calc(p)
	fmt.Print(a)
}

func calc(p Triangulo) int {
	return p.base * p.altura / 2
}
