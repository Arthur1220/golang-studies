package pagamento

import "fmt"

type CartaoCredito struct {
	NumeroCartao string
}

func (c CartaoCredito) Pagar(valor float64) string {
	return "Pagando R$ " + fmt.Sprintf("%.2f", valor) + " no cartão final " + c.NumeroCartao[len(c.NumeroCartao)-4:]
}
