package command

import "github.com/urfave/cli/v2"

func DeleteCommand() *cli.Command {
	return &cli.Command{
		Name:            "add",
		HelpName:        "add",
		Action:          deleteAction,
		Usage:           `create a http request`,
		Description:     `It's a curl package where you can send simple http requests. You can get your output as a file.`,
		SkipFlagParsing: true,
		HideHelp:        true,
		HideHelpCommand: true,
	}
}

func deleteAction(c *cli.Context) error {

	return nil
}
