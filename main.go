package main

import (
	"fmt"
	"os"
)

// ai_documentation_writer - Generate docs from code using AI
func ai_documentation_writer(path string) {
	fmt.Println("========================================")
	fmt.Println("  AI-Documentation-Writer")
	fmt.Println("  Generate docs from code using AI")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	ai_documentation_writer(path)
}
