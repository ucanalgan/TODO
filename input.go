package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInt(prompt string) (int, error) {
	for {
		input := readLine(prompt)
		value, err := strconv.Atoi(input)
		if err == nil {
			return value, nil
		}
		fmt.Println("Invalid input. Please enter a valid integer.")
	}
}
