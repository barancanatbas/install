package create

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/barancanatbas/install/model"
	"github.com/urfave/cli/v2"
)

const (
	INIT_FILE_NAME = "install.dependency.json"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:            "create",
		HelpName:        "create",
		Action:          createAction,
		Usage:           `create a http request`,
		Description:     `It's a curl package where you can send simple http requests. You can get your output as a file.`,
		SkipFlagParsing: true,
		HideHelp:        true,
		HideHelpCommand: true,
	}
}

func createAction(c *cli.Context) error {
	exists := fileExists(INIT_FILE_NAME)
	if exists {
		return errors.New("create var")
	}

	dummyData := model.Dependency{
		PackageName:    "khanmux",
		AllowedPackage: "github.com/barancanatbas/khanmux",
	}

	dummyBody, _ := json.Marshal(dummyData)

	err := CreateFile(INIT_FILE_NAME, dummyBody)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
