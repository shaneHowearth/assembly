package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	ui()
}

func ui() {
	scanner := bufio.NewScanner(os.Stdin)
	// arguments can be inside brackets, or just spaces
	r1, _ := regexp.Compile("\\(| ")
	// r2, _ := regexp.Compile(`[)].*`)
	// Initial Prompt
	fmt.Println("Simple book buying app, Type 'help' for a list of commands.")
	fmt.Print("> ")
	// Infinite loop waiting for input
	for scanner.Scan() {
		c := r1.Split(scanner.Text(), -1)
		// Anything after the command is deemed an argument
		//  var arg string
		//  if len(c) > 1 {
		//  	arg = string(r2.ReplaceAll([]byte(strings.Join(c[1:], " ")), []byte("")))
		//  }
		switch com := strings.ToLower(c[0]); com {
		case "help":
			Help()
		case "exit":
			os.Exit(0)
		}
		fmt.Print("> ")
	}

	if scanner.Err() != nil {
		os.Exit(1)
	}
}

// Help -
func Help() {
	fmt.Println(`
Available commands:
   Help - Prints this message.
   Exit - Exits this program.`)

}
