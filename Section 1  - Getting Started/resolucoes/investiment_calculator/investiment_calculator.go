package main

import "fmt"

//LEGENDA: lucroSemImposto: LSI, lucroComImposto: LCI

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculosFinancas(receita, despesas, taxaImposto float64) (float64, float64, float64) {
	lsi := receita - despesas
	lci := lsi - (lsi * taxaImposto / 100)
	razao := lci / receita
	return lsi, lci, razao
}


func main() {
	var receita, despesas, taxaImposto float64
	
	//Entradas do usuário
	receita = getUserInput("Digite o valor da receita: ")
	despesas = getUserInput("Digite o valor das despesas: ")
	taxaImposto = getUserInput("Digite a taxa de imposto: ")

	lsi, lci, razao := calculosFinancas(receita, despesas, taxaImposto)
	
	// O Sptint é uma função que retorna uma string formatada

	// Criando strings formatadas
	lsiStr := fmt.Sprintf("Lucro sem imposto: %.2f", lsi)
	lciStr := fmt.Sprintf("Lucro com imposto: %.2f", lci)
	razaoStr := fmt.Sprintf("Razão lucro/receita: %.2f", razao)

	// Exibindo as strings formatadas
	text := `Informações do Investimento:`
	fmt.Println(text)
	fmt.Println(lsiStr)
	fmt.Println(lciStr)
	fmt.Println(razaoStr)

	/*
	// Exibindo os resultados formatados
	fmt.Printf("Lucro sem imposto: %.2f\n", lsi)
	fmt.Printf("Lucro com imposto: %.2f\n", lci)
	fmt.Printf("Razão lucro/receita: %.2f\n", razao)
	*/
}