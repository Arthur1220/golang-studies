package contas

import "POO/05/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (conta *ContaPoupanca) Sacar(valor float64) bool {
	if valor > 0 && conta.saldo >= valor {
		conta.saldo -= valor
		return true
	}
	return false
}

func (conta *ContaPoupanca) Depositar(valor float64) bool {
	if valor > 0 {
		conta.saldo += valor
		return true
	}
	return false
}

func (conta *ContaPoupanca) Transferencia(valor float64, contaDestino *ContaPoupanca) bool {
	if valor > 0 && conta.saldo >= valor {
		conta.saldo -= valor
		contaDestino.saldo += valor
		return true
	}
	return false
}

func (conta *ContaPoupanca) GetSaldo() float64 {
	return conta.saldo
}
