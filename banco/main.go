package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func main() {
	contaDoTuts := ContaCorrente{"Arthur", 589, 123456, 125.30}
	fmt.Println(contaDoTuts)

	//var contaDaCris *ContaCorrente
	//contaDaCris = new(ContaCorrente)
	//contaDaCris.titular = "Cris"

	//fmt.Println(contaDaCris)

	fmt.Println(contaDoTuts.Sacar(-100.))
	fmt.Println(contaDoTuts.Depositar(-300.))

}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!\nSaldo: R$", c.saldo
	} else {
		return "Saldo insuficiente! Saldo:\nR$", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
	} else {
		return "Não é possivel depositar valores negativos!\nSaldo: R$", c.saldo
	}

	return "Deposito feito com sucesso! Saldo:\nR$", c.saldo
}
