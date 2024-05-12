package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	gameName := strings.Join(args, " ")
	if len(gameName) == 0 {
		fmt.Println("You need to provide a game")
		os.Exit(1)
	}
	gs := NewGameScrap(gameName)
	gs.Search()
}
