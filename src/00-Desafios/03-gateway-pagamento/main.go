/*
### 🟠 Nível 3: O Desafio da Interface (Sistema de Pagamento)

**Foco:** Interfaces, Structs e Polimorfismo.

**Cenário:**
Você está construindo um e-commerce. Você precisa processar pagamentos via **Pix** e **Cartão de Crédito**.

1.  Crie uma interface `ProcessadorPagamento` com o método `Pagar(valor float64) string`.
2.  Crie a struct `Pix` (campo: ChavePix) e `CartaoCredito` (campo: NumeroCartao).
3.  Implemente o método `Pagar` para ambos.
      * Pix retorna: "Pagando R$ X via Pix na chave Y".
      * Cartão retorna: "Pagando R$ X no cartão final Z (pegue os ultimos 4 digitos)".
4.  No `main`, crie uma lista (slice) de pagamentos misturados e processe todos com um loop.

**Dica de Go:** Lembre-se que para uma struct "ser" uma interface, basta implementar o método. Não precisa de `implements`.

*/

package main

import (
	"04-Desafios/03-gateway-pagamento/pagamento"
	"fmt"
)

func main() {
	fmt.Println("Desafio do Gateway de Pagamento!")

	// Aqui você criaria instâncias de Pix e CartaoCredito,
	// armazenaria em um slice de ProcessadorPagamento
	// e iteraria para chamar o método Pagar de cada um.

	pagamentos := []pagamento.ProcessadorPagamento{
		pagamento.Pix{ChavePix: "123456789"},
		pagamento.CartaoCredito{NumeroCartao: "1234567890123456"},
		pagamento.CartaoCredito{NumeroCartao: "9876543210987654"},
		pagamento.Pix{ChavePix: "987654321"},
		pagamento.CartaoCredito{NumeroCartao: "1111222233334444"},
		pagamento.Pix{ChavePix: "555555555"},
		pagamento.CartaoCredito{NumeroCartao: "9999888877776666"},
	}

	for _, pagamento := range pagamentos {
		fmt.Println(pagamento.Pagar(100.0))
	}
}
