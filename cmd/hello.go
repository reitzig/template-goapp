package cmd

import (
	"github.com/reitzig/template-goapp/internal"
	"github.com/spf13/cobra"
	"log"
)

var (
	helloConfig = internal.HelloConfig{}
)

func init() {
	Cmd.Flags().BoolVar(&helloConfig.Dear.Value, "dear", false, "say hello to a friend")
}

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "say hello to the world",
	Long:  `say hello to the world, or a friend!`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		log.Println(internal.Hello(helloConfig, args[0]))
	},
}
