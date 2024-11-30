package clicommands

import (
	"4devs-scrapping/src/Cli/Scrapers/CNPJ"
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

func CreateGeneratorCNPJCommand() cli.Command {
	return cli.Command{
		Name: "generate:cnpj",
		Aliases: []string{"-gcnpj"},
		Usage: "Generate a valid CNPJ",
		Action: func (ctx *cli.Context) error {
			fmt.Println(CNPJ.Generate(false))
			return nil
		},
	}
}

func CreateValidatorCNPJCommand() cli.Command {
	return cli.Command{
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
	}
}