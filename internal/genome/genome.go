package genome

import (
	"math/rand"

	"github.com/MaxHalford/eaopt"
)

const (

	// InputSize represents the number of inputs into the neural network.
	// Each berry/pest takes 5 inputs so that they can be one-hot encoded. This results in 25 inputs for the jars
	// Each category takes one input to represent whether or not it has been used.
	// Three inputs are necessary to one-hot encode how many rolls are left.
	// 25 (jars) + 9 (categories) + 3 (rolls left) = 37.
	InputSize = 37

	// OutputSize represents the number of outputs from the neural network.
	// Each category requires one output representing whether or not it should be scored.
	// There is one additional output representing the option to roll again.
	// There are also 5 outputs representing which jars to lock, if the output is to re-roll.
	// 9 (categories) + 1 (re-roll) + 5 (jars) = 14
	OutputSize = 14
)

// Genome is an object the implements the eaopt.Genome interface.
// It holds the weights and biases of a neural network.
type Genome struct {
	// Biases is a two-dimensional array representing the biases of the neural network.
	Biases [][]float64 `json:"biases"`

	// HiddenLayerSizes is an array of integers representing the size of hidden layers.
	// This array should be empty if there are no hidden layers.
	HiddenLayerSizes []int `json:"hiddenLayerSizes"`

	// Weights is a three-dimensional array of float64 representing the weights of the neural network.
	Weights [][][]float64 `json:"weights"`
}

func (g *Genome) Evaluate() (float64, error) {
	return 0.0, nil
}

func (g *Genome) Mutate(rng *rand.Rand) {

}

func (g *Genome) Crossover(genome eaopt.Genome, rng *rand.Rand) {
	
}

func (g *Genome) Clone() eaopt.Genome {
	copyG := &Genome{
		HiddenLayerSizes: append([]int{}, g.HiddenLayerSizes...), // Copy HiddenLayerSizes
	}

	// Deep copy Biases
	copyG.Biases = make([][]float64, len(g.Biases))
	for i := range g.Biases {
		copyG.Biases[i] = append([]float64{}, g.Biases[i]...)
	}

	// Deep copy Weights
	copyG.Weights = make([][][]float64, len(g.Weights))
	for i := range g.Weights {
		copyG.Weights[i] = make([][]float64, len(g.Weights[i]))
		for j := range g.Weights[i] {
			copyG.Weights[i][j] = append([]float64{}, g.Weights[i][j]...)
		}
	}

	return copyG
}