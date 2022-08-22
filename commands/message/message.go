package message

import (
	"flag"
	"fmt"
)

type Message struct {
	Text  string
	Helpf bool
}

const helpTextStart = `Responsible for print a simple text`
const helpLongTextStart = `
Responsible for print a simple text.

go_message_cli message --text [printText]
- printText: string
`
const exampleTextStart = `
go_message_cli message --text hello_world
`

func (cmd *Message) Name() string {
	return "message"
}

func (cmd *Message) Example() string {
	return exampleTextStart
}

func (cmd *Message) Help() string {
	return helpTextStart
}

func (cmd *Message) LongHelp() string {
	return helpLongTextStart
}

func (cmd *Message) Register(fs *flag.FlagSet) {
	fs.StringVar(&cmd.Text, "text", "", "printed text")
	fs.BoolVar(&cmd.Helpf, "help", false, "show help")
}

func (cmd *Message) Run() {
	if cmd.Helpf {
		fmt.Println(cmd.LongHelp())
		return
	}

	if cmd.Text == "" {
		fmt.Println("[--text] is required")
		return
	}

	fmt.Printf("Text received: %s\n", cmd.Text)
}
