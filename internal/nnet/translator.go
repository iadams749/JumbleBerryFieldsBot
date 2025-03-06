package nnet

import (
	"fmt"
	"sort"

	"github.com/iadams749/JumbleBerryFieldsBot/internal/game"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

// TranslateGameState takes a game state and converts it to an input tensor for the neural net.
func TranslateGameState(gs *game.GameState) *tensor.Dense {
	var inputs []float64

	if len(gs.Jars) != 5 {
		panic("invalid jar length")
	}

	// making the berry inputs
	for _, jar := range gs.Jars {
		switch jar.Berry {
		case game.Jumbleberry:
			inputs = append(inputs, []float64{1.0, 0.0, 0.0, 0.0, 0.0}...)
		case game.Sugarberry:
			inputs = append(inputs, []float64{0.0, 1.0, 0.0, 0.0, 0.0}...)
		case game.Pickleberry:
			inputs = append(inputs, []float64{0.0, 0.0, 1.0, 0.0, 0.0}...)
		case game.Moonberry:
			inputs = append(inputs, []float64{0.0, 0.0, 0.0, 1.0, 0.0}...)
		case game.Pest:
			inputs = append(inputs, []float64{0.0, 0.0, 0.0, 0.0, 1.0}...)
		default:
			panic("unrecognized berry")
		}
	}

	// encoding the round inputs
	switch gs.RollsLeftInTurn {
	case 2:
		inputs = append(inputs, []float64{1.0, 0.0, 0.0}...)
	case 1:
		inputs = append(inputs, []float64{0.0, 1.0, 0.0}...)
	case 0:
		inputs = append(inputs, []float64{0.0, 0.0, 1.0}...)
	default:
		panic("must have rolled at least once")
	}

	// encoding the categories
	if gs.Categories.JumbleberryCategory.Used {
		inputs = append(inputs, 1.0)
	} else {
		inputs = append(inputs, 0.0)
	}

	if gs.Categories.SugarberryCategory.Used {
		inputs = append(inputs, 1.0)
	} else {
		inputs = append(inputs, 0.0)
	}

	if gs.Categories.PickleberryCategory.Used {
		inputs = append(inputs, 1.0)
	} else {
		inputs = append(inputs, 0.0)
	}

	if gs.Categories.MoonberryCategory.Used {
		inputs = append(inputs, 1.0)
	} else {
		inputs = append(inputs, 0.0)
	}

	if gs.Categories.ThreeCategory.Used {
		inputs = append(inputs, 1.0)
	} else {
		inputs = append(inputs, 0.0)
	}

	if gs.Categories.FourCategory.Used {
		inputs = append(inputs, 1.0)
	} else {
		inputs = append(inputs, 0.0)
	}

	if gs.Categories.FiveCategory.Used {
		inputs = append(inputs, 1.0)
	} else {
		inputs = append(inputs, 0.0)
	}

	if gs.Categories.MixedCategory.Used {
		inputs = append(inputs, 1.0)
	} else {
		inputs = append(inputs, 0.0)
	}

	if gs.Categories.FreeCategory.Used {
		inputs = append(inputs, 1.0)
	} else {
		inputs = append(inputs, 0.0)
	}

	// returning the input as a tensor
	return tensor.New(
		tensor.WithBacking(inputs),
		tensor.WithShape(1, 37), // Shape matches the input layer
	)
}

func DoMoveFromTensor(game *game.GameState, output *gorgonia.Node) error {
	outputs, topIndices, err := GetTopKValues(output, 15)
	if err != nil {
		return fmt.Errorf("DoMoveFromTensor: %w", err)
	}

outerLoop:
	for _, val := range topIndices {
		switch val {
		case 0, 1, 2, 3, 4:
			continue
		case 5:
			if !game.Categories.JumbleberryCategory.Used {
				game.ScoreCategory(game.Categories.JumbleberryCategory)
				break outerLoop
			}
		case 6:
			if !game.Categories.SugarberryCategory.Used {
				game.ScoreCategory(game.Categories.SugarberryCategory)
				break outerLoop
			}
		case 7:
			if !game.Categories.PickleberryCategory.Used {
				game.ScoreCategory(game.Categories.PickleberryCategory)
				break outerLoop
			}
		case 8:
			if !game.Categories.MoonberryCategory.Used {
				game.ScoreCategory(game.Categories.MoonberryCategory)
				break outerLoop
			}
		case 9:
			if !game.Categories.ThreeCategory.Used {
				game.ScoreCategory(game.Categories.ThreeCategory)
				break outerLoop
			}
		case 10:
			if !game.Categories.FourCategory.Used {
				game.ScoreCategory(game.Categories.FourCategory)
				break outerLoop
			}
		case 11:
			if !game.Categories.FiveCategory.Used {
				game.ScoreCategory(game.Categories.FiveCategory)
				break outerLoop
			}
		case 12:
			if !game.Categories.MixedCategory.Used {
				game.ScoreCategory(game.Categories.MixedCategory)
				break outerLoop
			}
		case 13:
			if !game.Categories.FreeCategory.Used {
				game.ScoreCategory(game.Categories.FreeCategory)
				break outerLoop
			}
		case 14:
			if game.RollsLeftInTurn > 0 {
				// iterating over the output again to see which jars to lock or unlock
				for idx, val := range topIndices {
					if val > 4 {
						continue
					} else {
						// locking or unlocking the jars based on the output
						if outputs[idx] > 0 {
							game.Jars[val].Lock()
						} else {
							game.Jars[val].Unlock()
						}
					}
				}
				// rolling the jars
				game.RollJars()

				break outerLoop
			}

		}
	}

	return nil
}

// GetTopKValues takes a *gorgonia.Node and returns the top K values and their indices
// This is used to determine the desired action by examining the values in the final layer
func GetTopKValues(node *gorgonia.Node, k int) ([]float64, []int, error) {
	// Extract the tensor value from the node
	tensorVal := node.Value()

	// Ensure the node contains a tensor of a type we can work with
	dense, ok := tensorVal.(*tensor.Dense)
	if !ok {
		return nil, nil, fmt.Errorf("expected a *tensor.Dense, got %T", tensorVal)
	}

	// Flatten the tensor into a 1D slice of floats (if it's not already)
	data := dense.Data().([]float64)

	// Create a list of indices and values
	type pair struct {
		value float64
		index int
	}
	var values []pair
	for i, v := range data {
		values = append(values, pair{value: v, index: i})
	}

	// Sort by value (descending)
	sort.Slice(values, func(i, j int) bool {
		return values[i].value > values[j].value
	})

	// Get the top K values and indices
	var topKValues []float64
	var topKIndices []int
	for i := 0; i < k && i < len(values); i++ {
		topKValues = append(topKValues, values[i].value)
		topKIndices = append(topKIndices, values[i].index)
	}

	return topKValues, topKIndices, nil
}
