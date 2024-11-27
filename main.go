package main

import (
	CPF "4devs-scrapping/src/Cpf"
	"fmt"
)

func main() {
	cpf := CPF.Generate(true, "SP")
	fmt.Print("O CPF é: ", cpf)
}