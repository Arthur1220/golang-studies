/*
### 🟢 Nível 1: O Clássico do LeetCode (Two Sum)

**Foco:** Slices, Loops (For) e If/Else.

**Problema:**
Dado um array de inteiros `nums` e um inteiro `target`, retorne os **índices** dos dois números que somados resultam no `target`.
*Assuma que cada entrada terá exatamente uma solução e você não pode usar o mesmo elemento duas vezes.*

**Exemplo:**

```go
Input: nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
// Explicação: nums[0] + nums[1] == 9, então retornamos [0, 1].
```

*/

package main

import "fmt"

func main() {

	var nums = []int{2, 7, 11, 15}
	var target = 9

	result, check := twoSum(nums, target)

	if check {
		fmt.Println("Resultado encontrado:")
		fmt.Printf("nums[%d] + nums[%d] == %d, então retornamos [%d, %d].\n", nums[result[0]], nums[result[1]], target, result[0], result[1])
	} else {
		fmt.Println("Nenhum par encontrado.")
	}

}

func twoSum(nums []int, target int) ([]int, bool) {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}, true
			}
		}
	}

	return nil, false
}
