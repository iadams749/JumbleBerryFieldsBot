package genome

import (
	"github.com/iadams749/JumbleBerryFieldsBot/internal/game"
	"gorgonia.org/gorgonia"
)

func PlayGameFromGraph(g *gorgonia.ExprGraph, input *gorgonia.Node, output *gorgonia.Node) int {
	gs := game.NewGame()
	gs.RollJars()

	for gs.RoundsCompleted < 9 {
		// Create VM to run computation
		vm := gorgonia.NewTapeMachine(g)
		defer vm.Close()

		// Generating the input from the game state
		i := TranslateGameState(gs)

		// Assign input data to the input node
		if err := gorgonia.Let(input, i); err != nil {
			panic(err.Error())
		}

		// Run the computation graph
		if err := vm.RunAll(); err != nil {
			panic(err.Error())
		}

		// Doing the move based off of the output
		if err := DoMoveFromTensor(gs, output); err != nil {
			panic(err.Error())
		}
	}

	return gs.Score
}
