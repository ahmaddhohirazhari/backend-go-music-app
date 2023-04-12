package command

import (
	"backend/src/configs/server"
	database "backend/src/database/gorm"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple api with golang",
}

func init() {
	initCommand.AddCommand(database.MigrateCmd)
	initCommand.AddCommand(server.ServeCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
