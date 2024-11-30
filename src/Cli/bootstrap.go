package cli

import (
	"4devs-scrapping/src/cli/commands"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func Start() {
	cliApp := cli.NewApp()
	cliApp.Name = "4Devs Cli Scrapper"
  	cliApp.Usage = "Scraping the most usaged features from 4Devs"
	cliApp.Commands = []cli.Command{
		commands.CreateGeneratorCPFCommand(),
		commands.CreateGeneratorCNPJCommand(),
		commands.CreateGeneratorPersonCommand(),
		commands.CreateValidatorCPFCommand(),
		commands.CreateValidatorCNPJCommand(),
	}
	cliApp.ExitErrHandler = func(context *cli.Context, err error) {
		if err != nil {
			fmt.Println(err)
			cli.OsExiter(1)
		}
	}
	cliApp.Run(os.Args)
}