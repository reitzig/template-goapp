package main

import (
	"fmt"
	"github.com/reitzig/template-goapp/cmd"
	"os"
)

// goreleaser injected values
var version = "dev"

func main() {
	cmd.Cmd.Version = "v" + version

	err := cmd.Cmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(77)
	}
}
