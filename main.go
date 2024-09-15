package main

import (
	"bufio"
	"evaluator"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter expression:")

	for scanner.Scan() {
		line := scanner.Text()

		if line == "exit" {
			break
		}

		res, err := evaluator.Evaluate(line)
		if err != nil {
			panic(err)
			return
		}
		fmt.Printf("Result: %v\n", res)
		fmt.Println("Enter expression:")
	}
}
