package internal

import "regexp"

type UserCommand struct {
	Command   string
	Arguments []string
}

func argumentFilter(commandList []string) UserCommand {
	regexValidator := regexp.MustCompile("(?m)-")

	commandSet := false

	var commandDefinition UserCommand

	for _, argument := range commandList {
		isMatch := regexValidator.MatchString(argument)

		if !isMatch && !commandSet {
			commandDefinition.Command = argument
			commandSet = true
		} else if isMatch {
			commandDefinition.Arguments = append(commandDefinition.Arguments, argument)
		}
	}

	return commandDefinition
}
