package command

import (
	"backend/src/configs/server"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple api with golang",
}

func init() {
	initCommand.AddCommand(server.ServeCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}