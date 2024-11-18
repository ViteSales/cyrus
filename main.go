package main

import (
	"github.com/vitesales/cyrus/cmd"
	"log"
)

func main() {
	err := cmd.CyrusCmd.Execute()
	if err != nil {
		log.Fatal("Error executing command", "error", err)
	}
}
