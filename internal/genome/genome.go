package genome

import (
	"fmt"
	"math/rand"

	"github.com/MaxHalford/eaopt"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
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
	// 9 (categories) + 1 (re-roll) + 5 (jars) = 15
	OutputSize = 15
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
	graph, input, output, err := g.BuildGraph()
	if err != nil {
		return 0.0, err
	}

	score := PlayGameFromGraph(graph, input, output)
	fitness := 237.0 - float64(score)
	
	return fitness, nil
}

// Mutate applies random Gaussian noise to weights and biases to simulate mutation.
func (g *Genome) Mutate(rng *rand.Rand) {
	const mutationRate = 0.1     // Probability of each weight/bias being mutated
	const mutationStrength = 0.5 // Standard deviation of Gaussian noise

	// Mutate biases
	for i := range g.Biases {
		for j := range g.Biases[i] {
			if rng.Float64() < mutationRate {
				g.Biases[i][j] += rng.NormFloat64() * mutationStrength
			}
		}
	}

	// Mutate weights
	for i := range g.Weights {
		for j := range g.Weights[i] {
			for k := range g.Weights[i][j] {
				if rng.Float64() < mutationRate {
					g.Weights[i][j][k] += rng.NormFloat64() * mutationStrength
				}
			}
		}
	}
}

// Crossover performs uniform crossover between two Genomes.
func (g *Genome) Crossover(other eaopt.Genome, rng *rand.Rand) {
	otherGenome, ok := other.(*Genome)
	if !ok {
		panic("Cannot cast eaopt.Genome as *Genome")
	}

	// Crossover biases
	for i := range g.Biases {
		for j := range g.Biases[i] {
			if rng.Float64() < 0.5 {
				g.Biases[i][j], otherGenome.Biases[i][j] = otherGenome.Biases[i][j], g.Biases[i][j]
			}
		}
	}

	// Crossover weights
	for i := range g.Weights {
		for j := range g.Weights[i] {
			for k := range g.Weights[i][j] {
				if rng.Float64() < 0.5 {
					g.Weights[i][j][k], otherGenome.Weights[i][j][k] = otherGenome.Weights[i][j][k], g.Weights[i][j][k]
				}
			}
		}
	}
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

// BuildGraph builds a Gorgonia computation graph from the genome.
// It returns the graph, the input node, and the final output node.
func (g *Genome) BuildGraph() (graph *gorgonia.ExprGraph, input *gorgonia.Node, output *gorgonia.Node, err error) {
	graph = gorgonia.NewGraph()

	// Determine input and output sizes
	if len(g.Weights) == 0 || len(g.Weights[0]) == 0 {
		err = fmt.Errorf("invalid genome: weights not defined")
		return
	}

	// Create input node
	input = gorgonia.NewMatrix(graph,
		tensor.Float64,
		gorgonia.WithShape(1, InputSize), // batch size 1
		gorgonia.WithName("input"),
		gorgonia.WithInit(gorgonia.Zeroes()),
	)

	x := input

	// Build each layer
	for i := range g.Weights {
		wShape := tensor.Shape{len(g.Weights[i][0]), len(g.Weights[i]), }
		bShape := tensor.Shape{1, len(g.Biases[i])}

		// Create weight node
		wVal := tensor.New(tensor.WithShape(wShape...), tensor.WithBacking(flatten2D(g.Weights[i])))
		w := gorgonia.NewMatrix(graph,
			tensor.Float64,
			gorgonia.WithShape(wShape...),
			gorgonia.WithName(fmt.Sprintf("W%d", i)),
			gorgonia.WithValue(wVal),
		)

		// Create bias node
		bVal := tensor.New(tensor.WithShape(bShape...), tensor.WithBacking(g.Biases[i]))
		b := gorgonia.NewMatrix(graph,
			tensor.Float64,
			gorgonia.WithShape(bShape...),
			gorgonia.WithName(fmt.Sprintf("B%d", i)),
			gorgonia.WithValue(bVal),
		)

		// x = x * W + B
		var wx *gorgonia.Node
		if wx, err = gorgonia.Mul(x, w); err != nil {
			panic(err.Error())
		}

		var z *gorgonia.Node
		if z, err = gorgonia.Add(wx, b); err != nil {
			panic(err.Error())
		}

		// Apply activation (ReLU for all hidden layers, no activation for last layer)
		if i < len(g.Weights)-1 {
			if x, err = gorgonia.Rectify(z); err != nil {
				panic(err.Error())
			}
		} else {
			x = z // Final layer, no activation (or add softmax here if needed)
		}
	}

	output = x
	return
}

func flatten2D(matrix [][]float64) []float64 {
	flat := make([]float64, 0, len(matrix)*len(matrix[0]))
	for _, row := range matrix {
		flat = append(flat, row...)
	}
	return flat
}