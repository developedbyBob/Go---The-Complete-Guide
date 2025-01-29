package main

import "fmt"
import "math"



func calculadoraTest() {
	/*
	var investimentoInicial float64 = 1000.0
	taxaJuros := 5.5
	var anos float64 = 5 //anos aqui é float64 para poder ser usado na função math.Pow

	// ou investimentoInicial, taxaJuros, anos := 1000.0, 5.5, 5 -  o ponto flutuante ( .0 ) assegura que Go os tratará como float64

 	//var valorFuturo := investimentoInicial * math.Pow(1 + (taxaJuros/100), anos) vai da erro se o var tiver na frente

	valorFuturo := investimentoInicial * math.Pow(1 + (taxaJuros/100), anos) // estou usando :=


	fmt.Println("O valor futuro do investimento é de: ", valorFuturo)
	*/

	//outra forma de fazer

	var investimentoInicial float64
	var anos int
	var taxaJurosEsperada float64
	const inflacao float64 = 2.5
	
	fmt.Print("Digite o valor do investimento inicial: ")
	fmt.Scan(&investimentoInicial)

	fmt.Print("Digite a quantidade de anos: ")
	fmt.Scan(&anos)

	fmt.Print("Digite a taxa de juros esperada: ")
	fmt.Scan(&taxaJurosEsperada)

	//Calculando o valor futuro do investimento
	valorFuturo := investimentoInicial * math.Pow(1 + (taxaJurosEsperada/100), float64(anos))

	//Calculando o valor futuro do investimento com a inflação
	valorFuturoInflacao := valorFuturo / math.Pow(1 + (inflacao/100), float64(anos))

	fmt.Printf("==============Detalhes do Investimento===================\n")
	fmt.Printf("Investimento Inicial: %.2f\n", investimentoInicial)
	fmt.Printf("Quantidade de Anos: %d\n", anos)
	fmt.Printf("Taxa de Juros Esperada: %.2f\n", taxaJurosEsperada)
	fmt.Printf("Inflação: %.2f\n", inflacao)
	fmt.Println("")
	fmt.Printf("O valor futuro do investimento é de: %.2f\n", valorFuturo)
	fmt.Printf("O valor futuro do investimento com a inflação é de: %.2f\n", valorFuturoInflacao)
}
