package runserver

import (
	"github.com/agung96tm/miblog/bootstrap"
	"github.com/agung96tm/miblog/lib"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var configFile string

func init() {
	pf := StartCmd.PersistentFlags()
	pf.StringVarP(&configFile, "config", "c", "./config/config.yaml", "this parameter is used to start the service application")
}

var StartCmd = &cobra.Command{
	Use:          "runserver",
	Short:        "Start Server",
	SilenceUsage: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		lib.SetConfigPath(configFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func runApplication() {
	fx.New(bootstrap.Module, fx.NopLogger).Run()
}
