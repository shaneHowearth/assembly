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
	r2, _ := regexp.Compile(`[)].*`)
	// Initial Prompt
	fmt.Println("Simple book buying app, Type 'help' for a list of commands.")
	fmt.Print("> ")
	// Infinite loop waiting for input
	for scanner.Scan() {
		c := r1.Split(scanner.Text(), -1)
		switch com := strings.ToLower(c[0]); com {
		case "exit":
			os.Exit(0)
		}
		fmt.Print("> ")
	}

	if scanner.Err() != nil {
		os.Exit(1)
	}
}
