package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		exit(args)
	case "echo":
		echo(args)
	case "type":
		type_(args[0])
	default:
		if found, _ := doesFileExist(command); found {
			runProgram(command, args)
		} else {
			fmt.Fprint(os.Stdout, command+": command not found\n")
		}

	}

}

func doesFileExist(fileName string) (bool, string) {

	var found bool
	var result string
	paths := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	for _, path := range paths {
		fp := filepath.Join(path, fileName)
		_, err := os.Stat(fp)
		if err == nil {
			found = true
			result = fp
			break
		}
	}
	return found, result
}

func runProgram(fileName string, args []string) {
	cmd := exec.Command(fileName, args...)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
		return
	}
	fmt.Fprint(os.Stdout, string(stdout))
	return

}

func type_(comm string) {

	if slices.Contains(BUILTINS, comm) {
		fmt.Fprint(os.Stdout, comm+" is a shell builtin\n")
		return
	}
	if found, path_ := doesFileExist(comm); found {
		fmt.Fprint(os.Stdout, comm+" is ", path_)
	} else {
		fmt.Fprint(os.Stdout, comm+": not found")
	}
	fmt.Fprint(os.Stdout, "\n")
}

func echo(args []string) {
	for _, s := range args {
		fmt.Fprint(os.Stdout, s+" ")
	}
	fmt.Fprint(os.Stdout, "\n")
}

func exit(args []string) {

	if len(args) < 1 {
		fmt.Fprint(os.Stdout, "Not enough arguments!")
		os.Exit(1)
	}
	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprint(os.Stdout, err)
		os.Exit(1)
	}

	os.Exit(exitCode)
}
