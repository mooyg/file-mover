package main

import (
	"log"
	"os"

	"github.com/mooyg/file-mover/mover"
)

func main() {
	m, err := mover.NewFileMover("/Users/mooy/Desktop/test")

	if err != nil {
		log.Fatal("some error occured")
		os.Exit(0)
	}
	m.MoveFiles()
}
