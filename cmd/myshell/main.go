package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var BUILTINS = []string{"echo", "type", "exit"}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		inputParts := strings.Split(input, " ")
		command := inputParts[0]
		args := inputParts[1:]
		handleCommand(command, args)

	}
}

func handleCommand(command string, args []string) {
	command = strings.TrimSpace(command)
	for i, s := range args {
		args[i] = strings.TrimSpace(s)
	}

	switch command {
	case "exit":
		if len(args) < 1 {
			fmt.Println("Not enough arguments!")
			os.Exit(1)
		}
		exitCode, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Exit(exitCode)
	case "echo":
		for _, s := range args {
			fmt.Print(s + " ")
		}
		fmt.Print("\n")
	case "type":
		if slices.Contains(BUILTINS, args[0]) {
			fmt.Println(args[0] + " is a shell builtin")
		} else {
			fmt.Println(args[0] + ": command not found")
		}

	default:
		fmt.Println(command + ": command not found")

	}

}
