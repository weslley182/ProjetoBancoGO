package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	contaDaCris := new(ContaCorrente)
	contaDaCris.titular = "Cristiane"
	contaDaCris.saldo = 5000.95

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDaCris)
}
