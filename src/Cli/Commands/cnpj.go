package commands

import (
	"4devs-scrapping/src/cli/scrapers/cnpj"
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
			fmt.Println(cnpj.Generate(false))
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
			cnpjArg := ctx.Args().First()
			if cnpjArg == "" {
				return errors.New("o cnpj deve ser informado")
			}
			fmt.Println(cnpj.Validator(cnpjArg))
			return nil
		},
	}
}