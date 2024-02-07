// Para resolver o problema de encontrar o comprimento da maior substring
// sem caracteres repetidos em Go, podemos usar uma abordagem semelhante à
// descrita anteriormente, adaptada para Go. Utilizaremos um mapa para
// rastrear os índices mais recentes dos caracteres já vistos e duas
// variáveis para manter o início da janela atual (`start`) e o comprimento
// máximo da substring encontrada (`maxLength`).

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	start, maxLength := 0, 0
	lastSeen := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if index, found := lastSeen[s[i]]; found && index >= start {
			start = index + 1
		}
		lastSeen[s[i]] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
}

// Nesta solução, iteramos sobre a string `s` uma vez, usando um índice `i`
// para representar a extremidade direita da janela atual. O mapa
// `lastSeen` mantém os índices dos últimos locais em que cada caractere
// foi visto. Se encontrarmos um caractere que já foi visto e o índice
// armazenado no mapa for igual ou superior ao início da janela atual
// (`start`), atualizamos `start` para o índice seguinte à última
// ocorrência desse caractere. Isso efetivamente "pula" a janela para
// frente, excluindo o caractere repetido e qualquer caractere antes dele
// que poderia estar na janela anterior. A cada passo, atualizamos o
// comprimento máximo da substring (`maxLength`) se a largura da janela
// atual (calculada como `i - start + 1`) for maior que o `maxLength` atual.
// Esta abordagem garante que cada caracter seja considerado exatamente uma
// vez, resultando em uma complexidade de tempo de \(O(n)\), onde \(n\) é o
// comprimento da string `s`. A complexidade de espaço é \(O(min(m, n))\),
// onde \(m\) é o tamanho do conjunto de caracteres no alfabeto (por
// exemplo, o número total de caracteres ASCII se a string contiver apenas
// caracteres ASCII), pois o mapa `lastSeen` armazena no máximo um índice
// para cada caractere único presente na string `s`.
