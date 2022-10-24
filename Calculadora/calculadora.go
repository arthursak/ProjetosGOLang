package main

import (
	"fmt"
)

var num1 int
var num2 int
var opc int

func main() {
	menu()

	for opc != 0 {
		fmt.Println("Digite o primeiro número: ")
		fmt.Scan(&num1)
		fmt.Println("Digite o segundo número: ")
		fmt.Scan(&num2)

		switch opc {
		case 1:
			somar()

		case 2:
			subtracao()

		case 3:
			multiplicacao()

		case 4:
			divisao()

		}

		menu()
	}

	fmt.Println("Até mais!")

}

func menu() {
	fmt.Println("------------------------")
	fmt.Println("O que você deseja realizar?\n1 - Soma\n2 - Subtração\n3 - Multiplicação\n4 - Divisão\n0 - Sair")
	fmt.Println("------------------------")
	fmt.Scan(&opc)
}

func somar() int {
	rst := num1 + num2
	soma, _ := fmt.Println("A soma entre os números", num1, "e", num2, "é:", rst)
	return soma
}

func subtracao() int {
	rst := num1 - num2
	sub, _ := fmt.Println("A subtração entre os números", num1, "e", num2, "é:", rst)
	return sub
}

func multiplicacao() int {
	rst := num1 * num2
	multi, _ := fmt.Println("A multiplacação entre os números", num1, "e", num2, "é:", rst)
	return multi
}

func divisao() int {
	rst := num1 / num2
	div, _ := fmt.Println("A divisão entre os números", num1, "e", num2, "é:", rst)
	return div
}
