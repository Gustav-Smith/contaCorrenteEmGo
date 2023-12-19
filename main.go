package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoGustavo := contas.ContaCorrente{}
	contaDoGustavo.Depositar(500)
	PagarBoleto(&contaDoGustavo, 6000)

	fmt.Println(contaDoGustavo.ObterSaldo())
}
