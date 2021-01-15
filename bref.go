package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jph5396/bref/commands"
)

func main() {
	config := commands.CommandConfig{}
	global := flag.NewFlagSet("global", flag.ExitOnError)
	// NOTE: print not implemented yet.
	// global.BoolVar(&config.Print, "print", false, "print results of scrape.")
	global.StringVar(&config.SaveDir, "usedir", "", "path to an alternate directory to save files. needs to already exist.")

	comms := commands.NewCommandFactory()
	global.Usage = func() {
		printHelp(comms)
		fmt.Println("global options:")
		global.PrintDefaults()
	}
	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arugments.")
		global.Usage()
		os.Exit(1)
	}
	parseErr := global.Parse(os.Args[3:])
	if parseErr != nil {
		fmt.Printf("parse error: %v \n", parseErr.Error())
	}
	c, ok := comms.Get(os.Args[1])
	if ok {
		c.SetConfig(config)
		err := c.Run(os.Args[2:])
		if err != nil {
			fmt.Printf("Error: %v. use flag --help for command docs. \n", err.Error())
		}

	} else {
		fmt.Println("Error: Command not found.")
		global.Usage()
	}

}

func printHelp(c commands.CommandFactory) {
	c.PrintCommands()
}
