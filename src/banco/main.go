package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoTuts := contas.ContaCorrente{Titular: "Arthur", Saldo: 200.50}
	fmt.Println(contaDoTuts)

	contaDaFran := contas.ContaCorrente{Titular: "Francine", Saldo: 200.50}
	fmt.Println(contaDaFran)

	status := contaDoTuts.Transferir(100., &contaDaFran)

	fmt.Println(status)
	fmt.Println(contaDoTuts)
	fmt.Println(contaDaFran)

}
