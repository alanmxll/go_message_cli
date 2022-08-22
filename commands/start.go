package commands

import (
	"flag"
	"fmt"
	"net/http"
)

type Start struct {
	Port    int
	Version string
	Helpf   bool
}

const helpTextStart = `Responsible for starting a simple server`
const helpLongTextStart = `
Responsible for starting a simple server releasing access to all
configured services.

go_message_cli start --port [serverPort] --version [serverVersion]
- serverPort: int
- servverVersion: string
`
const exampleTextStart = `
go_message_cli start --port 3030 --version 1.0.0
go_message_cli start --version 1.0.0
`

func (cmd *Start) Name() string {
	return "start"
}

func (cmd *Start) Example() string {
	return exampleTextStart
}

func (cmd *Start) Help() string {
	return helpTextStart
}

func (cmd *Start) LongHelp() string {
	return helpLongTextStart
}

func (cmd *Start) Register(fs *flag.FlagSet) {
	fs.IntVar(&cmd.Port, "port", 8080, "port to listen on")
	fs.StringVar(&cmd.Version, "version", "", "version of the server")
	fs.BoolVar(&cmd.Helpf, "help", false, "show help")
}

func (cmd *Start) Run() {
	if cmd.Helpf {
		fmt.Println(cmd.LongHelp())
		return
	}

	if cmd.Version == "" {
		fmt.Println("[--version] is required")
		return
	}

	fmt.Printf("Starting server on port %d with version %s\n", cmd.Port, cmd.Version)
	http.ListenAndServe(fmt.Sprintf(":%d", cmd.Port), nil)
}
