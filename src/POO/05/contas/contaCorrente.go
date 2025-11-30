package contas

import "POO/05/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Sacar(valor float64) bool {
	if valor > 0 && conta.saldo >= valor {
		conta.saldo -= valor
		return true
	}
	return false
}

func (conta *ContaCorrente) Depositar(valor float64) bool {
	if valor > 0 {
		conta.saldo += valor
		return true
	}
	return false
}

func (conta *ContaCorrente) Transferencia(valor float64, contaDestino *ContaCorrente) bool {
	if valor > 0 && conta.saldo >= valor {
		conta.saldo -= valor
		contaDestino.saldo += valor
		return true
	}
	return false
}

func (conta *ContaCorrente) GetSaldo() float64 {
	return conta.saldo
}
