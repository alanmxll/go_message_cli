package main

import (
	"fmt"

	"github.com/alanmxll/go_message_cli/commands"
	"github.com/alanmxll/go_message_cli/commands/message"
	"github.com/alanmxll/go_message_cli/internal"
)

var commandList = []internal.Command{
	new(commands.Start),
	new(message.Message),
}

func main() {
	err := internal.CommandInit("go_message_cli").Start(commandList)

	if err != nil {
		fmt.Println(err.Error())
	}
}
