package main

import (
	"Banco/clientes"
	"Banco/contas"
	"fmt"
)

func main() {
	silvia := clientes.Titular{Nome: "Silvia"}
	gustavo := clientes.Titular{Nome: "Gustavo"}

	contaDaSilvia := contas.ContaCorrente{Titular: silvia}
	contaDaSilvia.Depositar(300)

	contaDoGustavo := contas.ContaCorrente{Titular: gustavo}
	contaDoGustavo.Depositar(100)

	//status := contaDaSilvia.Tranferir(100, &contaDoGustavo)
	status := contaDoGustavo.Tranferir(100, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia.ObterSaldo())
	fmt.Println(contaDoGustavo.ObterSaldo())

	contas.PagarBoleto(&contaDaSilvia, 60)
	fmt.Println(contaDaSilvia.ObterSaldo())
}
