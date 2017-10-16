package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/vapor-ware/ksync/pkg/cli"
)

var (
	globalUsage = `Map container names to local filesystem locations.`

	RootCmd = &cobra.Command{
		Use:   "radar",
		Short: "Map container names to local filesystem locations.",
		Long:  globalUsage,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cli.InitLogging()
		},
	}
)

// Main runs the server instance
func main() {
	RootCmd.AddCommand(
		(&serveCmd{}).new(),
	)

	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}

// Init initializes the server instance
func init() {
	cobra.OnInitialize(func() { cli.InitConfig("radar") })

	cli.DefaultFlags(RootCmd, "radar")
}
