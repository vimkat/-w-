package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"kitty"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("[FATAL] Expected one file name to load\n")
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName) // should end in .mew
	if err != nil {
		log.Fatalf("[FATAL] Failed to read file %s: %s\n", fileName, err)
	}

	code, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("[FATAL] Failed to read code from file: %s\n", err)
	}
	file.Close()

	kitty := kitty.New(string(code)) // code to run
	for !kitty.IsDead() {
		kitty.HuntLaser()

		if kitty.IsObserved() {
			fmt.Printf("%s\n", kitty.CurrentStacc())
			fmt.Printf("%s\n", kitty.CurrentInstruction())
		}
	}
}
