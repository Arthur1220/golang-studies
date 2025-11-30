package contas

// IConta define os métodos obrigatórios para qualquer tipo de conta bancária.
// O 'I' na frente não é obrigatório em Go, mas ajuda na didática aqui.
// Em Go, interfaces geralmente terminam em 'er' (ex: Payer, Reader), mas usaremos Conta para clareza.
type Conta interface {
	Sacar(valor float64) bool
	Depositar(valor float64) bool
	GetSaldo() float64
}
