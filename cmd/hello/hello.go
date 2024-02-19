package hello

import (
	"github.com/reitzig/template-goapp/internal"
	"github.com/spf13/cobra"
	"log"
)

var (
	dear bool
)

func init() {
	Cmd.Flags().BoolVar(&dear, "dear", false, "say hello to a friend")
}

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "say hello to the world",
	Long:  `say hello to the world, or a friend!`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(internal.Hello(dear, args[0]))
	},
}
