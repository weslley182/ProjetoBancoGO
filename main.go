package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	contaDaCris := new(ContaCorrente)
	contaDaCris.titular = "Cristiane"
	contaDaCris.saldo = 5000.95

	contaDaCris2 := new(ContaCorrente)
	contaDaCris2.titular = "Cristiane"
	contaDaCris2.saldo = 5000.95

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	fmt.Println(contaDaCris == contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)

	fmt.Println(contaDaCris.saldo)

	fmt.Println(contaDaCris.Sacar(400))
	fmt.Println(contaDaCris.saldo)
}
