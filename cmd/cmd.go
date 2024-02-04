package cmd

import (
	"errors"
	"github.com/agung96tm/miblog/cmd/makemigrations"
	"github.com/agung96tm/miblog/cmd/migrate"
	"github.com/agung96tm/miblog/cmd/runserver"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(runserver.StartCmd)
	rootCmd.AddCommand(makemigrations.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
}

var rootCmd = &cobra.Command{
	Use:   "miblog",
	Short: "micro blog",
	Long:  "micro blog created with golang programming language",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run:              func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
