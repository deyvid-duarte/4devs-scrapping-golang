package main

import (
	"4devs-scrapping/src/CNPJ"
	"4devs-scrapping/src/CPF"
	"4devs-scrapping/src/Person"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "4Devs Cli Scrapper"
  	cliApp.Usage = "Scraping the most usaged features from 4Devs"
	cliApp.Commands = []cli.Command{
		{
			Name: "generate:cpf",
			Aliases: []string{"-gcpf"},
			Usage: "Generate a valid CPF",
			Action: func (ctx *cli.Context) error {
				fmt.Println(CPF.Generate(true, ""))
				return nil
			},
		},
		{
			Name: "generate:cnpj",
			Aliases: []string{"-gcnpj"},
			Usage: "Generate a valid CNPJ",
			Action: func (ctx *cli.Context) error {
				fmt.Println(CNPJ.Generate(false))
				return nil
			},
		},
		{
			Name: "generate:person",
			Aliases: []string{"-gperson"},
			Usage: "Generate a random person",
			Action: func (ctx *cli.Context) error {
				person := Person.Person{Gender: "M", Age: 32}
				fmt.Println(person.Generate(false, 1))
				return nil
			},
		},
		{
			Name: "validate:cpf",
			Aliases: []string{"-vcpf"},
			Usage: "Validate CPF",
			Action: func (ctx *cli.Context) error {
				cpf := ctx.Args().First()
				if cpf == "" {
					return errors.New("o cpf deve ser informado")
				}
				fmt.Println(CPF.Validator(cpf))
				return nil
			},
		},
		{
			Name: "validate:cnpj",
			Aliases: []string{"-vcnpj"},
			Usage: "Validate CNPJ",
			Action: func (ctx *cli.Context) error {
				cnpj := ctx.Args().First()
				if cnpj == "" {
					return errors.New("o cnpj deve ser informado")
				}
				fmt.Println(CNPJ.Validator(cnpj))
				return nil
			},
		},
	}
	cliApp.ExitErrHandler = func(context *cli.Context, err error) {
		if err != nil {
			fmt.Println(err)
			cli.OsExiter(1)
		}
	}
	cliApp.Run(os.Args)
}