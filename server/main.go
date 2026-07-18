package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter text (Press Ctrl+C to stop):")

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Captured: %s\n", line)
	}

}
