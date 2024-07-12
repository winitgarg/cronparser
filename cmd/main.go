package main

import (
	"fmt"
	"os"

	"github.com/cronparser/internal/parser"
)

func main() {
	// Usage information
	usage := "Usage: ./cronparser \"*/15 0 1,15 * 1-5 /usr/bin/find\""

	if len(os.Args) < 2 {
		fmt.Println(usage)

		os.Exit(1)
	}

	cronExpr := os.Args[1]

	cronParser := parser.New()

	err := cronParser.Parse(cronExpr)
	if err != nil {
		fmt.Printf("error in parsing input: %s, err: %s", cronExpr, err.Error())
	}

	return
}
