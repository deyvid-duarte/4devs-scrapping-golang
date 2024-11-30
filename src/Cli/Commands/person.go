package clicommands

import (
	"4devs-scrapping/src/Cli/Scrapers/Person"
	"fmt"

	"github.com/urfave/cli"
)

func CreateGeneratorPersonCommand() cli.Command {
	return cli.Command{
		Name: "generate:person",
		Aliases: []string{"-gperson"},
		Usage: "Generate a random person",
		Action: func (ctx *cli.Context) error {
			person := Person.Person{Gender: "M", Age: 32}
			fmt.Println(person.Generate(false, 1))
			return nil
		},
	}
}