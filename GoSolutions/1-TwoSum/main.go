package main

// Para resolver este problema com uma complexidade de tempo melhor do que \(O(n^2)\),
// podemos usar um mapa (dicionário) para armazenar os índices dos elementos enquanto
// iteramos pelo array. Isso nos permite verificar se o complemento de cada elemento
// (isto é, o valor que, quando somado ao elemento, resulta no alvo) já foi visto no
// array com apenas uma passagem pelo array, alcançando assim uma complexidade de tempo
// \(O(n)\).

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// Mapa para armazenar os valores e seus índices
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		// Verifica se o complemento existe no mapa
		if index, found := numMap[complement]; found {
			// Retorna o índice do complemento e o índice atual
			return []int{index, i}
		}
		// Armazena o número atual e seu índice no mapa
		numMap[num] = i
	}
	// Retorna um slice vazio se não houver solução
	return []int{}
}

func main() {
	// Teste 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println("Test 1:", twoSum(nums1, target1)) // [0, 1]

	// Teste 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println("Test 2:", twoSum(nums2, target2)) // [1, 2]

	// Teste 3
	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println("Test 3:", twoSum(nums3, target3)) // [0, 1]
}

// Esta solução itera uma vez pelo array `nums`, usando um mapa para rastrear os
// índices dos elementos. Para cada elemento, calcula-se seu complemento (o valor
// necessário para que a soma com o elemento atual seja igual ao alvo) e verifica-se se
// esse complemento já foi encontrado anteriormente no array. Se o complemento foi
// encontrado, a função retorna os índices do elemento atual e do complemento. Caso
// contrário, o elemento atual é adicionado ao mapa, e a iteração continua. Esta
// abordagem garante que encontraremos a solução com complexidade de tempo linear
// \(O(n)\) e complexidade de espaço \(O(n)\), uma vez que podemos ter que armazenar
// todos os elementos no mapa na pior das hipóteses.
