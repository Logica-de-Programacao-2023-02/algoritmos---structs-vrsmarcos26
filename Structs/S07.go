//Crie uma struct chamada Animal com os campos "nome", "espécie", "idade" e "som".
//Escreva funções que permitam modificar o som que o animal faz e uma função que
//imprima as informações do animal e o som que ele faz.

package main

import "fmt"

type Animal struct {
	nome    string
	especie string
	idade   string
	som     string
}

var nomev string
var especiev string
var idadev string
var somv string
var som2 string

func modificadordevoz(a *Animal) {
	a.som = som2
}

func main() {
	fmt.Print("\nDiga qual seu animal: ")
	fmt.Scan(&nomev)
	fmt.Print("\nDiga qual especie do seu animal: ")
	fmt.Scan(&especiev)
	fmt.Print("\nDiga quantos anos tem seu animal: ")
	fmt.Scan(&idadev)
	fmt.Print("\nDiga qual som seu animal faz: ")
	fmt.Scan(&somv)
	fmt.Print("\nQue som deseje que ele faça: ")
	fmt.Scan(&som2)
	fmt.Print("\nA ficha do seu animal ERA ASSIM : ")
	fmt.Print(Animal{nome: nomev, especie: especiev, idade: idadev, som: somv})
	novoanimal := Animal{nome: nomev, especie: especiev, idade: idadev, som: som2}
	modificadordevoz(&novoanimal)
	fmt.Print("\nA ficha do seu animal FICOU ASSIM : ")
	fmt.Print(novoanimal)
}
