package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//检测输出
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	//读取用户输入直到遇到\n
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("Hello, %s! What can I do for you?\n", name)
	}

	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("an error occured: %s\n", err)
			continue
		}

		input = input[:len(input)-1]

		input = strings.ToLower(input)

		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry，I didn't catch you.")
		}
	}
}
