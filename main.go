// Package main is the main entry point for the groqmit command.
package main

import (
	_ "embed"
	"fmt"

	"github.com/conneroisu/groqmit/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
