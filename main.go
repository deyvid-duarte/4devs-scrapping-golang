package main

import (
	"4devs-scrapping/src/CNPJ"
	"4devs-scrapping/src/CPF"
	"fmt"
)

func main() {
	cpf := CPF.Generate(false, "SP")
	cnpj := CNPJ.Generate(false)
	fmt.Println("O CPF é:", cpf)
	fmt.Println("O CNPJ é:", cnpj)
}