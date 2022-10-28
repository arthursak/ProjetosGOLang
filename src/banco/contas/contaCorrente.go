package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!\nsaldo: R$", c.saldo
	} else {
		return "saldo insuficiente! saldo:\nR$", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransf float64, contaDestino *ContaCorrente) string {

	if valorDaTransf < c.saldo && valorDaTransf > 0 {
		c.saldo -= valorDaTransf
		contaDestino.Depositar(valorDaTransf)
		return "Transferencia efetuada com sucesso!"
	} else {
		return "Não há saldo suficiente para realizar a transferencia"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
	} else {
		return "Não é possivel depositar valores negativos!\nsaldo: R$", c.saldo
	}

	return "Deposito feito com sucesso! saldo:\nR$", c.saldo
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
