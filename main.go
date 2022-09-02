package main

import (
	"os"

	"github.com/amazemind-io/learning-paths-validator/validator"
)

func main() {
	argsWithoutProg := os.Args[1:]
	cmd := argsWithoutProg[0]
	path_file := argsWithoutProg[1]

	validator.ValidateCmd(cmd, path_file, argsWithoutProg)
}
