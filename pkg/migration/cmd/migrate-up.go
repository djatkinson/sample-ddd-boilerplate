package cmd

import (
	"ddd-boilerplate/config"
	"ddd-boilerplate/pkg/migration"
	"ddd-boilerplate/pkg/postgres"
	"fmt"

	"github.com/spf13/cobra"
)

var migrateUpCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use: "up",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up")

			config := config.NewConfig().PostgreSQLConfig
			db := postgres.ConnectDBForMigration(config)
			migration.PostgresMigrate(db)
		},
	}

	migrateCmd.AddCommand(migrateUpCmd)
}
