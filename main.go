package main

import (
	"fmt"
	"math/rand"

	"github.com/MaxHalford/eaopt"
	"github.com/iadams749/JumbleBerryFieldsBot/internal/nnet"
)

func main() {
	// Instantiate a GA with a GAConfig
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	ga.NGenerations = 1000
	ga.PopSize = 1000
	ga.ParallelEval = true
	ga.Model = eaopt.ModGenerational{
		Selector:  eaopt.SelTournament{NContestants: 3},
		MutRate:   0.2,
		CrossRate: 0.7,
	}

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	// Initialize the GA with the neural network factory
	if err := ga.Minimize(func(rng *rand.Rand) eaopt.Genome {
		return nnet.NewNeuralNetPlayer(32, 32)
	}); err != nil {
		panic(err)
	}
}
