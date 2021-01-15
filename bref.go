package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jph5396/bref/commands"
)

func main() {
	gameData := commands.GetGame(os.Args[1])

	err := JSONFileWriter("test.json", gameData)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// JSONFileWriter reusable function for writing structs to json files.
// will return any errors that occur, or nil if it succeeds.
func JSONFileWriter(path string, data interface{}) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	encodeErr := encoder.Encode(data)
	if encodeErr != nil {
		return encodeErr
	}

	defer file.Close()
	fmt.Println("Created file: ", path)
	return nil
}
