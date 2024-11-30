package clicommands

import (
	"4devs-scrapping/src/Cli/Scrapers/CPF"
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

func CreateGeneratorCPFCommand() cli.Command {
	return cli.Command{
		Name: "generate:cpf",
		Aliases: []string{"-gcpf"},
		Usage: "Generate a valid CPF",
		Action: func (ctx *cli.Context) error {
			fmt.Println(CPF.Generate(true, ""))
			return nil
		},
	}
}

func CreateValidatorCPFCommand() cli.Command {
	return cli.Command{
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
	}
}