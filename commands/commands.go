//Package commands defines the commands got bref and the factory used by the main command to
// to retrieve a command for execution or gather help command information.
package commands

//Command basic interface that should be implemented by a command
type Command interface {
	//CommandName should return the name of the command.
	CommandName() string
	//Parse should accept the command line args that are parsed with the
	// flag package. any error that occurs should be returned.
	Parse([]string) error
	//Run executes actual command. it should return a json encoded byte array if any
	// data is to be saved, and any error that occurs or nil if everything was successful.
	Run() ([]byte, error)
}
