package cmd

import (
	"github.com/reitzig/template-goapp/cmd/hello"
	"github.com/spf13/cobra"
	"os"
)

var (
	Debug bool
)

func init() {
	Cmd.PersistentFlags().BoolVar(&Debug, "debug", false, "debug output")

	Cmd.AddCommand(hello.Cmd)
}

var Cmd = &cobra.Command{
	Use:   "template-goapp",
	Short: "a stub",
	Long:  `a stub app to try out some stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(1)
	},
}
