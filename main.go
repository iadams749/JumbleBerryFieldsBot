package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/iadams749/JumbleBerryFieldsBot/internal/genome"
)

func main() {
	// // Instantiate a GA with a GAConfig
	// var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// ga.NGenerations = 1000
	// ga.PopSize = 100
	// ga.ParallelEval = true
	// ga.Model = eaopt.ModGenerational{
	// 	Selector:  eaopt.SelTournament{NContestants: 3},
	// 	MutRate:   0.2,
	// 	CrossRate: 0.7,
	// }

	// // Add a custom print function to track progress
	// ga.Callback = func(ga *eaopt.GA) {
	// 	fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	// }

	// // Initialize the GA with the neural network factory
	// if err := ga.Minimize(func(rng *rand.Rand) eaopt.Genome {
	// 	return nnet.NewNeuralNetPlayer(32, 32)
	// }); err != nil {
	// 	panic(err)
	// }

	src := rand.NewSource(0)
	r := rand.New(src)

	hiddenLayers := []int{2,4,8}

	genome := genome.NewGenome(r, hiddenLayers)
	// Convert struct to pretty JSON
	jsonData, err := json.MarshalIndent(genome, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Define the filename
	filename := "output.json"

	// Write JSON data to the file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data successfully written to", filename)
}
