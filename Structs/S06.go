package main

import "fmt"

type Funcionario struct {
	nome    string
	salario int
	idade   int
}

func aumentoediminuicaodesalario(f *Funcionario) {
	if f.salario <= 1000 && f.idade > 18 {
		f.salario = int(float64(f.salario) * 1.5)
	}
}

func tempodeservico(f *Funcionario) {
	if f.idade < 18 {
		fmt.Print("Desculpe, você não tem idade para trabalhar, ou seja, não é um funcionário nosso.")
	} else {
		f.idade -= 18
	}
}

func main() {
	var nomev string
	var salariov int
	var idadev int

	fmt.Print("Qual seu nome: ")
	fmt.Scan(&nomev)
	fmt.Print("Quanto você ganha: ")
	fmt.Scan(&salariov)
	fmt.Print("Qual sua idade? ")
	fmt.Scan(&idadev)

	funcionario := Funcionario{nome: nomev, salario: salariov, idade: idadev}

	tempodeservico(&funcionario)
	aumentoediminuicaodesalario(&funcionario)

	fmt.Print("Após analisarmos sua ficha temos: ")
	fmt.Print(funcionario)
}
