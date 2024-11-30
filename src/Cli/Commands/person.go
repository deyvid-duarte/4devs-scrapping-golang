package commands

import (
	"4devs-scrapping/src/cli/scrapers/person"
	"fmt"

	"github.com/urfave/cli"
)

func CreateGeneratorPersonCommand() cli.Command {
	return cli.Command{
		Name: "generate:person",
		Aliases: []string{"-gperson"},
		Usage: "Generate a random person",
		Action: func (ctx *cli.Context) error {
			person := person.Person{Gender: "M", Age: 32}
			fmt.Println(person.Generate(false, 1))
			return nil
		},
	}
}