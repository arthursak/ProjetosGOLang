package main

import (
	"fmt"
)

var prod string
var opc int

func main() {
	menu()
	adicionaLista()
}

func menu() {
	fmt.Println("Quer adicionar algo na lista de produtos?")
	fmt.Scan(&opc)

	for opc != 0 {
		adicionaLista()
	}
}

func adicionaLista() {
	fmt.Println("Digite um produto que você gostaria de adicionar: ")
	fmt.Scan(&prod)
	var listaProd []string
	listaProd = append(listaProd, prod)

	fmt.Println(listaProd)
	fmt.Println("Quer adicionar algo na lista de produtos?")
	fmt.Scan(&opc)

}
