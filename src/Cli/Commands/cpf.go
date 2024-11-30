package commands

import (
	"4devs-scrapping/src/cli/scrapers/cpf"
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
			fmt.Println(cpf.Generate(true, ""))
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
			cpfArg := ctx.Args().First()
			if cpfArg == "" {
				return errors.New("o cpf deve ser informado")
			}
			fmt.Println(cpf.Validator(cpfArg))
			return nil
		},
	}
}