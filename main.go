package main

import (
	"fmt"

	"github.com/MaxHalford/eaopt"
	"github.com/iadams749/JumbleBerryFieldsBot/internal/genome"
)

func main() {
	// Instantiate a GA with a GAConfig
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	ga.NGenerations = 1000
	ga.PopSize = 100
	ga.ParallelEval = true
	ga.Model = eaopt.ModGenerational{
		Selector:  eaopt.SelTournament{NContestants: 3},
		MutRate:   0.2,
		CrossRate: 0.7,
	}

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		var totalFitness float64
        for _, indiv := range ga.Populations[0].Individuals {
            totalFitness += indiv.Fitness
        }
        avgFitness := totalFitness / float64(len(ga.Populations[0].Individuals))
        fmt.Printf("Generation %d | Avg Fitness: %f | Best: %f\n", ga.Generations, avgFitness, ga.HallOfFame[0].Fitness)
	}

	// Initialize the GA with the neural network factory
	if err := ga.Minimize(genome.NewGenomeFactory([]int{128, 128})); err != nil {
		panic(err)
	}
}
