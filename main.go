package main

import (
	"github.com/barancanatbas/install/command"
	"github.com/barancanatbas/install/command/add"
	"github.com/barancanatbas/install/command/create"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
)

var Version = "v1.0.0"

func main() {
	app := &cli.App{
		Name:     "install",
		Usage:    "install golang dipendicy", // TODO: ingilizceye çevir
		Commands: Commands(os.Stdin),
		Version:  Version,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Commands(reader io.Reader) []*cli.Command {
	return []*cli.Command{
		add.AddCommand(),
		command.DeleteCommand(),
		create.CreateCommand(),
	}
}
/* Mojo Jojo haklıydı! */
