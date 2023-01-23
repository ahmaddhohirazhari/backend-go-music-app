package command

import (
	"backend/src/configs/serve"
	database "backend/src/database/gorm"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple api with golang",
}

func init() {
	initCommand.AddCommand(serve.ServeCmd)
	initCommand.AddCommand(database.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
