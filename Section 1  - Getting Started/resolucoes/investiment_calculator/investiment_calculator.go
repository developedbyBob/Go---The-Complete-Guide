package main

import "fmt"

//LEGENDA: lucroSemImposto: LSI, lucroComImposto: LCI


func main() {
	var receita, despesas, taxaImposto float64
	
	//Entradas do usuário
	fmt.Print("Digite o valor da receita: ")
	fmt.Scan(&receita)

	fmt.Print("Digite o valor das despesas: ")
	fmt.Scan(&despesas)

	fmt.Print("Digite a taxa de imposto: ")
	fmt.Scan(&taxaImposto)

	//Calculos
	lsi := receita - despesas
	lci := lsi - (lsi * taxaImposto / 100)
	razao := lci / receita

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