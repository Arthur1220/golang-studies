package pagamento

import "fmt"

type Pix struct {
	ChavePix string
}

func (p Pix) Pagar(valor float64) string {
	return "Pagando R$ " + fmt.Sprintf("%.2f", valor) + " via Pix na chave " + p.ChavePix
}
