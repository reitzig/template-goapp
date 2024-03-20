package cmd

import (
	"github.com/reitzig/template-goapp/internal"
	"github.com/spf13/cobra"
	"os"
)

var (
	config internal.RootConfig = internal.RootConfig{
		Hello: helloConfig,
	}
)

func init() {
	RootCmd.PersistentFlags().BoolVar(&config.Debug.Value, "debug", false, "debug output")

	RootCmd.AddCommand(infoCmd)
	RootCmd.AddCommand(Cmd)
}

var RootCmd = &cobra.Command{
	Use:   "template-goapp",
	Short: "a stub",
	Long:  `a stub app to try out some stuff`,
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
		os.Exit(1)
	},
}
