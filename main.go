package main

import (
	"awesomeProject/kata"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./data/words.txt")
	if err != nil {
		panic(err)
	}

	game, err := kata.NewGame(f)

	if err != nil {
		panic(err)
	}

	sp := game.GetShortestPath("cat", "dog")
	fmt.Printf("Shortest transformation path between `cat` and `dog`: %+v\n", sp)
}


