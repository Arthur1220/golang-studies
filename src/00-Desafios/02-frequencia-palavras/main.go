/*
### 🟡 Nível 2: Manipulação de Strings e Maps

**Foco:** `map`, `range`, `strings`.

**Problema:**
Crie uma função que receba uma frase (string) e retorne um mapa (`map[string]int`) indicando a frequência de cada palavra.

**Exemplo:**

```go
Input: "go é legal e go é rápido"
Output: map[string]int{
    "go": 2,
    "é": 2,
    "legal": 1,
    "e": 1,
    "rápido": 1,
}
```

**Dica de Go:**

  * Você vai precisar do pacote `strings`.
  * Use `strings.Fields(frase)` para quebrar a frase em um *slice* de palavras (ele já ignora os espaços).

-----
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Desafio da Frequencia de Palavras:")

	frase := "go é legal e go é rápido"

	frequencia := contarFrequenciaPalavras(frase)

	fmt.Println(frequencia)
}

func contarFrequenciaPalavras(frase string) map[string]int {
	frequencia := make(map[string]int)

	for _, palavra := range strings.Fields(frase) {
		frequencia[palavra]++
	}

	return frequencia
}
