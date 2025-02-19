package main

import (
	"fmt"

	"github.com/iadams749/JumbleBerryFieldsBot/internal/game"
)

func main() {
	game := game.NewGame()

	game.RollJars()

	fmt.Println(game)

	game.ScoreCategory(game.Categories.FreeCategory)

	fmt.Println(game)
}
