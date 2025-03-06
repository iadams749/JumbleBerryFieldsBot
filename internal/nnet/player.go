package nnet

import (
	"fmt"

	"github.com/iadams749/JumbleBerryFieldsBot/internal/game"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

const (
	inputSize  = 37 // 25 nodes for berries + 3 nodes for rolls left + 9 nodes for the categories
	outputSize = 15 // 10 nodes to represent 10 choices (9 categories or roll again) and 5 berries to represent whether or not to lock jar
)

type NeuralNetPlayer struct {
	Graph    *gorgonia.ExprGraph
	Input    *gorgonia.Node
	Hidden1  *gorgonia.Node
	Hidden2  *gorgonia.Node
	Output   *gorgonia.Node
	Weights1 *gorgonia.Node
	Biases1  *gorgonia.Node
	Weights2 *gorgonia.Node
	Biases2  *gorgonia.Node
	Weights3 *gorgonia.Node
	Biases3  *gorgonia.Node
}

func NewNeuralNetPlayer(hiddenSize1, hiddenSize2 int) *NeuralNetPlayer {
	g := gorgonia.NewGraph()

	// Define network parameters with random initialization
	w1 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(inputSize, hiddenSize1), gorgonia.WithName("w1"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	b1 := gorgonia.NewVector(g, tensor.Float64, gorgonia.WithShape(hiddenSize1), gorgonia.WithName("b1"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))

	w2 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(hiddenSize1, hiddenSize2), gorgonia.WithName("w2"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	b2 := gorgonia.NewVector(g, tensor.Float64, gorgonia.WithShape(hiddenSize2), gorgonia.WithName("b2"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))

	w3 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(hiddenSize2, outputSize), gorgonia.WithName("w3"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	b3 := gorgonia.NewVector(g, tensor.Float64, gorgonia.WithShape(outputSize), gorgonia.WithName("b3"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))

	// Define input node
	input := gorgonia.NewTensor(g, tensor.Float64, 2, gorgonia.WithShape(1, inputSize), gorgonia.WithName("input"))

	// Forward propagation
	h1PreAct, _ := gorgonia.Add(gorgonia.Must(gorgonia.Mul(input, w1)), b1)
	h1 := gorgonia.Must(gorgonia.Tanh(h1PreAct))

	h2PreAct, _ := gorgonia.Add(gorgonia.Must(gorgonia.Mul(h1, w2)), b2)
	h2 := gorgonia.Must(gorgonia.Tanh(h2PreAct))

	output, _ := gorgonia.Add(gorgonia.Must(gorgonia.Mul(h2, w3)), b3)

	return &NeuralNetPlayer{
		Graph:    g,
		Input:    input,
		Hidden1:  h1,
		Hidden2:  h2,
		Output:   output,
		Weights1: w1,
		Biases1:  b1,
		Weights2: w2,
		Biases2:  b2,
		Weights3: w3,
		Biases3:  b3,
	}
}

func (n *NeuralNetPlayer) DoMoveFromGameState(gs *game.GameState) error {
	// Create VM to run computation
	vm := gorgonia.NewTapeMachine(n.Graph)
	defer vm.Close()

	// Generating the input from the game state
	input := TranslateGameState(gs)

	// Assign input data to the input node
	if err := gorgonia.Let(n.Input, input); err != nil {
		return fmt.Errorf("error setting input: %w", err)
	}

	// Run the computation graph
	if err := vm.RunAll(); err != nil {
		return fmt.Errorf("error calculating graph: %w", err)
	}

	// Doing the move based off of the output
	if err := DoMoveFromTensor(gs, n.Output); err != nil {
		return fmt.Errorf("error doing move: %w", err)
	}

	return nil
}

func DeepCopyNeuralNetPlayer(original *NeuralNetPlayer) *NeuralNetPlayer {
	// Create a new graph for the deep copy
	g := gorgonia.NewGraph()

	// Copy the weights and biases using WithValue
	w1 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(original.Weights1.Shape()[0], original.Weights1.Shape()[1]), gorgonia.WithName("w1"), gorgonia.WithShape(original.Weights1.Shape()...), gorgonia.WithValue(original.Weights1.Value()))
	b1 := gorgonia.NewVector(g, tensor.Float64, gorgonia.WithShape(original.Biases1.Shape()[0]), gorgonia.WithName("b1"), gorgonia.WithShape(original.Biases1.Shape()...), gorgonia.WithValue(original.Biases1.Value()))

	w2 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(original.Weights2.Shape()[0], original.Weights2.Shape()[1]), gorgonia.WithName("w2"), gorgonia.WithShape(original.Weights2.Shape()...), gorgonia.WithValue(original.Weights2.Value()))
	b2 := gorgonia.NewVector(g, tensor.Float64, gorgonia.WithShape(original.Biases2.Shape()[0]), gorgonia.WithName("b2"), gorgonia.WithShape(original.Biases2.Shape()...), gorgonia.WithValue(original.Biases2.Value()))

	w3 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(original.Weights3.Shape()[0], original.Weights3.Shape()[1]), gorgonia.WithName("w3"), gorgonia.WithShape(original.Weights3.Shape()...), gorgonia.WithValue(original.Weights3.Value()))
	b3 := gorgonia.NewVector(g, tensor.Float64, gorgonia.WithShape(original.Biases3.Shape()[0]), gorgonia.WithName("b3"), gorgonia.WithShape(original.Biases3.Shape()...), gorgonia.WithValue(original.Biases3.Value()))

	// Copy the input node
	input := gorgonia.NewTensor(g, tensor.Float64, 2, gorgonia.WithShape(original.Input.Shape()...), gorgonia.WithName("input"))

	// Perform forward propagation like in the original player
	h1PreAct, _ := gorgonia.Add(gorgonia.Must(gorgonia.Mul(input, w1)), b1)
	h1 := gorgonia.Must(gorgonia.Tanh(h1PreAct))

	h2PreAct, _ := gorgonia.Add(gorgonia.Must(gorgonia.Mul(h1, w2)), b2)
	h2 := gorgonia.Must(gorgonia.Tanh(h2PreAct))

	output, _ := gorgonia.Add(gorgonia.Must(gorgonia.Mul(h2, w3)), b3)

	// Return the new deep copy NeuralNetPlayer
	return &NeuralNetPlayer{
		Graph:    g,
		Input:    input,
		Hidden1:  h1,
		Hidden2:  h2,
		Output:   output,
		Weights1: w1,
		Biases1:  b1,
		Weights2: w2,
		Biases2:  b2,
		Weights3: w3,
		Biases3:  b3,
	}
}