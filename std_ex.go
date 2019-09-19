package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var name string
	fmt.Println("What's your name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("Hi!", name)
}
