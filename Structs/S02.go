//Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço".
//O campo "endereço" deve ser uma outra struct com os campos "rua", "número", "cidade" e "estado".
//Escreva uma função que receba uma Pessoa como parâmetro e imprima seu endereço completo.

package main

import "fmt"

type Pessoa struct {
	nome     string
	idade    string
	endereço Endereço
}
type Endereço struct {
	rua    string
	numero string
	cidade string
	estado string
}

func main() {
	p := Pessoa{nome: "Marocs", idade: "19", endereço: Endereço{rua: "Cruziero", numero: "1", cidade: "Brasília", estado: "DF"}}
	dcr := ec(p)
	fmt.Print("Sua descrição seria completo seria: ", dcr)
}

func ec(p Pessoa) string {
	return fmt.Sprintf("%s, %s, %s, %s", p.endereço.rua, p.endereço.numero, p.endereço.cidade, p.endereço.estado)
}
