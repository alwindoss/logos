package main

import (
	"fmt"
	"log"
	"os"

	"imkernel.dev/logos/internal/app"
)

func main() {
	f, err := os.Open("./data/en/kjv.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bible, err := app.ParseOSISFile(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parsed the %s Bible from the osis file successfully!\n", bible.Identifier)
}
