package main

import "github.com/mooyg/file-mover/mover"

func main() {
	m := mover.NewFileMover("~/Desktop/backup")

	m.MoveFiles()
}
