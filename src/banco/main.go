package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaTuts := contas.ContaCorrente{}
	contaTuts.Depositar(100)

	fmt.Println(contaTuts.ObterSaldo())

}
