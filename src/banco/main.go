package main

import (
	"fmt"

	"./contas/"
)

func main() {
	contaDoTuts := contas.ContaCorrente{"Arthur", 589, 123456, 125.30}
	fmt.Println(contaDoTuts)

	contaDaFran := contas.ContaCorrente{"Francine", 755, 654321, 200.30}
	fmt.Println(contaDaFran)

	status := contaDoTuts.Transferir(100., &contaDaFran)

	//var contaDaCris *ContaCorrente
	//contaDaCris = new(ContaCorrente)
	//contaDaCris.titular = "Cris"

	//fmt.Println(contaDaCris)

	fmt.Println(status)
	fmt.Println(contaDoTuts)
	fmt.Println(contaDaFran)

}
