package contas

type ContaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso!\nSaldo: R$", c.Saldo
	} else {
		return "Saldo insuficiente! Saldo:\nR$", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransf float64, contaDestino *ContaCorrente) string {

	if valorDaTransf < c.Saldo && valorDaTransf > 0 {
		c.Saldo -= valorDaTransf
		contaDestino.Depositar(valorDaTransf)
		return "Transferencia efetuada com sucesso!"
	} else {
		return "Não há saldo suficiente para realizar a transferencia"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.Saldo += valorDeposito
	} else {
		return "Não é possivel depositar valores negativos!\nSaldo: R$", c.Saldo
	}

	return "Deposito feito com sucesso! Saldo:\nR$", c.Saldo
}
