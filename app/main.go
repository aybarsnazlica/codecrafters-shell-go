package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Printf("%s: command not found\n", command)
		return
	}
}
