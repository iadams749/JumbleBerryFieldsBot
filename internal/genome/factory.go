package genome

import (
	"math"
	"math/rand"

	"github.com/MaxHalford/eaopt"
)

// A GenomeFactory is a function that initalizes random genomes.
// It is used for the Minimize() call when training using the eaopt package.
type GenomeFactory func(rng *rand.Rand) eaopt.Genome

// NewGenomeFactory creates a GenomeFactory that initalizes genomes with the provided hiddel layer size.
func NewGenomeFactory(hiddenLayerSizes []int) GenomeFactory {
	return func(rng *rand.Rand) eaopt.Genome {
		return NewGenome(rng, hiddenLayerSizes)
	}
}

// NewGenome creates a new genome with the provided hidden layer sizes.
// Values are initalized usign glorot initalization.
func NewGenome(rng *rand.Rand, hiddenLayerSizes []int) *Genome {
	// initalize the genome
	g := &Genome{}
	g.HiddenLayerSizes = hiddenLayerSizes

	// build a list of size for each layer, including input/hidden/output
	sizes := []int{InputSize}
	sizes = append(sizes, hiddenLayerSizes...)
	sizes = append(sizes, OutputSize)

	// initalize the arrays according to the hidden sizes
	for i := 1; i < len(sizes); i++ {
		// initalize the biases
		bias := make([]float64, sizes[i])
		for idx := range sizes[i] {
			bias[idx] = glorot(rng, InputSize, OutputSize)
		}
		g.Biases = append(g.Biases, bias)

		// initalize the weights
		weights := make([][]float64, sizes[i])
		for outerIdx := range sizes[i] {
			weight := make([]float64, sizes[i-1])
			for innerIdx := range sizes[i-1] {
				weight[innerIdx] = glorot(rng, InputSize, OutputSize)
			}
			weights[outerIdx] = weight
		}
		g.Weights = append(g.Weights, weights)
	}

	return g
}

// glorot initializes a float64 using the Glorot uniform distribution.
func glorot(rng *rand.Rand, fanIn, fanOut int) float64 {
	// Seed for randomness
	limit := math.Sqrt(6.0 / float64(fanIn+fanOut))
	return rng.Float64()*(2*limit) - limit // Scale to [-limit, limit]
}
