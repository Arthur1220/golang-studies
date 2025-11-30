package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (conta *ContaCorrente) Sacar(valor float64) bool {
	if valor > 0 && conta.Saldo >= valor {
		conta.Saldo -= valor
		return true
	}
	return false
}

func (conta *ContaCorrente) Depositar(valor float64) bool {
	if valor > 0 {
		conta.Saldo += valor
		return true
	}
	return false
}

func (conta *ContaCorrente) Transferencia(valor float64, contaDestino *ContaCorrente) bool {
	if valor > 0 && conta.Saldo >= valor {
		conta.Saldo -= valor
		contaDestino.Saldo += valor
		return true
	}
	return false
}
