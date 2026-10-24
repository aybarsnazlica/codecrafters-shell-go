package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)

		builtins := []string{"exit", "type", "echo"}

		if command == "exit" {
			break
		} else if strings.HasPrefix(command, "type") {
			content := strings.SplitN(command, " ", 2)[1]
			if slices.Contains(builtins, content) {
				fmt.Printf("%s is a shell builtin\n", content)
				continue
			}
			fmt.Printf("%s: not found\n", content)
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
			continue
		}
	}
}
