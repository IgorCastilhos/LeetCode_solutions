// Para encontrar a mediana de dois arrays ordenados com uma complexidade
// de tempo \(O(\log(m+n))\), podemos usar um algoritmo de busca binária. A
// ideia central é particionar os dois arrays em duas metades tais que a
// primeira metade contém os elementos menores e a segunda metade os
// elementos maiores. Se conseguirmos fazer isso, a mediana estará entre o
// maior elemento da primeira metade e o menor elemento da segunda metade.
// Aqui está um esboço de como podemos fazer isso:
// 1. **Escolha o Array Menor:** Para minimizar o número de comparações,
// sempre aplicamos a busca binária no array de menor tamanho.
// 2. **Partição dos Arrays:** Vamos chamar os dois arrays de `A` e `B`,
// onde `A` é o menor. Se `A` e `B` têm tamanhos `m` e `n`,
// respectivamente, queremos encontrar um índice `i` em `A` e um índice `j`
// em `B` tais que `i + j = (m + n + 1) / 2`. Isso garante que a partição
// deixe exatamente metade (ou metade mais um, se `m + n` for ímpar) dos
// elementos nos dois lados.
// 3. **Ajuste de Índices:** Usamos busca binária para encontrar o índice
// `i` em `A`. O índice `j` em `B` pode ser calculado diretamente usando a
// fórmula `j = (m + n + 1) / 2 - i`.
// 4. **Verificação das Condições de Partição:** Precisamos garantir que
// `A[i-1] <= B[j]` e `B[j-1] <= A[i]`. Se não for esse o caso, ajustamos
// o valor de `i` usando busca binária.
// 5. **Cálculo da Mediana:**
// - Se `m + n` é ímpar, a mediana é o maior entre `A[i-1]` e `B[j-1]`.
// - Se `m + n` é par, a mediana é a média entre o maior dos elementos à
// esquerda da partição e o menor dos elementos à direita da partição.

package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	m, n := len(A), len(B)
	if m > n { // Garantindo que A é o menor array
		A, B, m, n = B, A, n, m
	}

	imin, imax, halfLen := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i
		if i < m && B[j-1] > A[i] {
			// i é muito pequeno, deve ser aumentado
			imin = i + 1
		} else if i > 0 && A[i-1] > B[j] {
			// i é muito grande, deve ser diminuído
			imax = i - 1
		} else { // i é perfeito
			maxOfLeft, minOfRight := 0.0, 0.0
			if i == 0 {
				maxOfLeft = float64(B[j-1])
			} else if j == 0 {
				maxOfLeft = float64(A[i-1])
			} else {
				maxOfLeft = float64(max(A[i-1], B[j-1]))
			}

			if (m+n)%2 == 1 {
				return maxOfLeft
			}

			if i == m {
				minOfRight = float64(B[j])
			} else if j == n {
				minOfRight = float64(A[i])
			} else {
				minOfRight = float64(min(A[i], B[j]))
			}

			return (maxOfLeft + minOfRight) / 2.0
		}
	}

	return 0.0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))

	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

// Esta implementação segue a abordagem descrita acima, garantindo que a
// complexidade de tempo seja \(O(\log(\min(m,n)))\), onde \(m\) e \(n\)
// são os tamanhos dos arrays de entrada. Isso é alcançado aplicando a
// busca binária no array menor.
