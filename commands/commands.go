//Package commands defines the commands got bref and the factory used by the main command to
// to retrieve a command for execution or gather help command information.
package commands

import (
	"fmt"
)

//Command basic interface that should be implemented by a command
type Command interface {
	//CommandName should return the name of the command.
	CommandName() string
	//CommandDescription should return the description that will be called when
	// the --help argument is passed.
	CommandDescription() string
	//SetConfig set the command config to use when the command runs.
	SetConfig(CommandConfig)
	//Run executes actual command. Should accept any arguments that are supplied at the command
	//line that might not be parsed by the flag set.
	Run([]string) error
}

//CommandConfig TODO: explain
type CommandConfig struct {
	Print   bool
	SaveDir string
}

//CommandFactory contains a map with the command name as the key and the command as the value.
type CommandFactory struct {
	Commands map[string]Command
}

//NewCommandFactory creates a new commandfactory with all commands in it.
func NewCommandFactory() CommandFactory {
	command := make(map[string]Command)
	dc := NewDayCommand()
	command[dc.CommandName()] = &dc
	gc := NewGameCommand()
	command[gc.CommandName()] = &gc
	return CommandFactory{
		Commands: command,
	}
}

//Get tries to locate the requested command and return it. bool will be set to false if none exists.
func (cf CommandFactory) Get(comm string) (Command, bool) {
	command, ok := cf.Commands[comm]
	return command, ok
}

//PrintCommands can be called to print all available commands to the command line.
func (cf CommandFactory) PrintCommands() {
	fmt.Println("Available Commands:")
	for _, val := range cf.Commands {
		fmt.Println(val.CommandDescription())
	}
}
