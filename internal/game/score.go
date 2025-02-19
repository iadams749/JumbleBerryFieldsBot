package game

import "fmt"

func (gs *GameState) ScoreCategory(cat Category) error {
	if gs.RollsLeftInTurn > 2 {
		return fmt.Errorf("must have rolled once to score category")
	}

	score, err := cat.CalcScore(gs.GetBerries())
	if err != nil {
		return fmt.Errorf("error calculating score: %w", err)
	}

	gs.Score += score
	gs.NewTurn()

	return nil
}