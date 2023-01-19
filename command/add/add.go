package add

import (
	"encoding/json"
	"fmt"
	"github.com/barancanatbas/install/model"
	"github.com/urfave/cli/v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:            "add",
		HelpName:        "add",
		Action:          addAction,
		Usage:           `create a http request`,
		Description:     `It's a curl package where you can send simple http requests. You can get your output as a file.`,
		SkipFlagParsing: true,
		HideHelp:        true,
		HideHelpCommand: true,
	}
}

func addAction(c *cli.Context) error {
	var request editRequest
	if err := request.BindRequest(c.Args().Slice()); err != nil {
		fmt.Println(err)
		return err
	}

	dependency := model.Dependency{
		PackageName:    request.PackageName,
		AllowedPackage: request.AllowedPackage,
	}

	dependenciesBody, err := json.Marshal(dependency)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = writeFile(request.FileName, dependenciesBody)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
