package cli

import (
	clicommands "4devs-scrapping/src/Cli/Commands"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func Start() {
	cliApp := cli.NewApp()
	cliApp.Name = "4Devs Cli Scrapper"
  	cliApp.Usage = "Scraping the most usaged features from 4Devs"
	cliApp.Commands = []cli.Command{
		clicommands.CreateGeneratorCPFCommand(),
		clicommands.CreateGeneratorCNPJCommand(),
		clicommands.CreateGeneratorPersonCommand(),
		clicommands.CreateValidatorCPFCommand(),
		clicommands.CreateValidatorCNPJCommand(),
	}
	cliApp.ExitErrHandler = func(context *cli.Context, err error) {
		if err != nil {
			fmt.Println(err)
			cli.OsExiter(1)
		}
	}
	cliApp.Run(os.Args)
}