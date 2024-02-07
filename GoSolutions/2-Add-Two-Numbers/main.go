// Para resolver este problema, podemos simular o processo de soma como
// faríamos no papel, somando dígito a dígito, começando do dígito menos
// significativo (o primeiro nó das listas ligadas).
// Mantemos um acompanhamento de qualquer "carry" (valor de transporte) que
// ocorra quando a soma de dois dígitos excede 9. Vamos iterar através das
// duas listas ligadas simultaneamente, somando os valores correspondentes
// dos nós, mais o valor de transporte, se houver. Se uma das listas for mais
// longa que a outra, continuaremos a iteração até que todos os nós de ambas
// as listas sejam processados. Se ao final ainda houver um valor de
// transporte, adicionaremos um novo nó com esse valor ao resultado.
// Aqui está uma implementação em pseudocódigo, seguido de uma implementação
// em Go:
// Pseudocódigo:
// Inicialize uma nova lista ligada "resultHead" com um nó sentinela (nó
// fictício inicial). Defina "current" para apontar para o nó sentinela.
// Inicialize "carry" como 0.
// Enquanto qualquer uma das listas "l1" ou "l2" não for nula ou "carry" for
// diferente de 0:
//   - Calcule a soma dos valores atuais de "l1" e "l2", mais o "carry".
//   - Atualize "carry" para ser a soma dividida por 10 (o novo valor de
//
// transporte).
//   - Crie um novo nó com o valor de soma % 10 e adicione-o após o nó
//
// "current".
//   - Avance "l1", "l2" e "current" se eles não forem nulos.
//
// Retorne "resultHead.next" (pula o nó sentinela).
package main

import "fmt"

// Definição para um nó de lista ligada.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Função para adicionar dois números representados por listas ligadas.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Nó sentinela para evitar verificações de nulo no loop.
	resultHead := &ListNode{0, nil}
	current := resultHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{sum % 10, nil}
		current = current.Next
	}

	// O primeiro nó é um nó sentinela; retornamos o próximo.
	return resultHead.Next
}

// Função auxiliar para imprimir a lista ligada.
func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
		l = l.Next
	}
	fmt.Println()
}

func main() {
	// Exemplo 1
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	printList(result) // Deve imprimir: 7 -> 0 -> 8

	// Adicione mais exemplos conforme necessário...
}

// Esta solução itera uma vez através das duas listas, resultando em uma
// complexidade de tempo \(O(\max(n,m))\), onde \(n\) e \(m\) são os
// comprimentos das listas ligadas `l1` e `l2`, respectivamente. A
// complexidade de espaço é \(O(\max(n,m))\) para a nova lista ligada
// criada como resultado.
