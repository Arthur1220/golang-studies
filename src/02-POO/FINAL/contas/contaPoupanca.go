package contas

import "POO/FINAL/clientes"

// ContaPoupanca estrutura de dados para contas poupança.
type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

// Sacar realiza a retirada de valores.
func (c *ContaPoupanca) Sacar(valor float64) bool {
	podeSacar := valor > 0 && c.saldo >= valor
	if podeSacar {
		c.saldo -= valor
		return true
	}
	return false
}

// Depositar adiciona valores positivos.
func (c *ContaPoupanca) Depositar(valor float64) bool {
	if valor > 0 {
		c.saldo += valor
		return true
	}
	return false
}

// Transferencia permite enviar dinheiro até para Contas Correntes.
func (c *ContaPoupanca) Transferencia(valor float64, contaDestino Conta) bool {
	if c.Sacar(valor) {
		contaDestino.Depositar(valor)
		return true
	}
	return false
}

// GetSaldo retorna o valor atual.
func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
