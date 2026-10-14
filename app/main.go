package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("$ ")
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("%s: command not found\n", command)
}
