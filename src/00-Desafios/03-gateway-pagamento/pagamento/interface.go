package pagamento

// ProcessadorPagamento define o contrato para processadores de pagamento.
// Qualquer método de pagamento deve implementar o método Pagar.
// Em Go, interfaces geralmente terminam em 'er' (ex: Payer, Reader), mas usaremos ProcessadorPagamento para clareza.
type ProcessadorPagamento interface {
	Pagar(valor float64) string
}
