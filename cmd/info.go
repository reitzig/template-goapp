package cmd

import (
	"github.com/reitzig/template-goapp/internal"
	"github.com/spf13/cobra"
	"log"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "information about ",
	Long:  `Prints information about the tool and its configuration`,
	Args:  cobra.ExactArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		log.Println(internal.Info(config))
	},
}
