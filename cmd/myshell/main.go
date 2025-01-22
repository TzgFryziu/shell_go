package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	_, err := fmt.Fprint(os.Stdout, "$ ")
	if err != nil {
		return
	}

	// Wait for user input
	_, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}
}
