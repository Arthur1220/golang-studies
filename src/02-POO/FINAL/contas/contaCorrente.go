package contas

import "POO/FINAL/clientes"

// ContaCorrente estrutura de dados para contas correntes.
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64 // saldo é privado (encapsulamento)
}

// Sacar realiza a retirada de valores se houver saldo.
func (c *ContaCorrente) Sacar(valor float64) bool {
	podeSacar := valor > 0 && c.saldo >= valor
	if podeSacar {
		c.saldo -= valor
		return true
	}
	return false
}

// Depositar adiciona valores positivos ao saldo.
func (c *ContaCorrente) Depositar(valor float64) bool {
	if valor > 0 {
		c.saldo += valor
		return true
	}
	return false
}

// Transferencia envia valor da conta atual para qualquer conta de destino.
// Note que contaDestino agora é do tipo 'Conta' (Interface), aceitando Poupança ou Corrente.
func (c *ContaCorrente) Transferencia(valor float64, contaDestino Conta) bool {
	if c.Sacar(valor) {
		contaDestino.Depositar(valor)
		return true
	}
	return false
}

// GetSaldo retorna o valor atual do saldo.
func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
