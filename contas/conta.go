package contas

import "fmt"

func PagarBoleto(conta Conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
	fmt.Println("Conta paga com sucesso.")
}

type Conta interface {
	Sacar(valor float64) string
}
