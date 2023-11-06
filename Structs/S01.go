//Crie uma struct chamada Circulo com o campo "raio". Escreva uma função que receba um Circulo como
//parâmetro e calcule a área do círculo (área = pi * raio * raio). Dica: use a constante math.
//Pi para representar o número pi.

package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func main() {
	p := Circulo{raio: 2}
	a := area(p)
	fmt.Print(a)
}

func area(p Circulo) float64 {
	return math.Pi * p.raio * p.raio
}
