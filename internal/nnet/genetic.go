package nnet

import (
	"errors"
	"math/rand"

	"github.com/MaxHalford/eaopt"
	"github.com/iadams749/JumbleBerryFieldsBot/internal/game"
	"gorgonia.org/gorgonia"
)

func (n *NeuralNetPlayer) Evaluate() (float64, error) {
	scores := make([]float64, 0)

	for range 100 {
		gs := game.NewGame()
		gs.RollJars()

		for gs.RoundsCompleted < 9 {
			n.DoMoveFromGameState(gs)
		}

		scores = append(scores, float64(gs.Score))
	}

	sum := 0.0

	for _, score := range scores {
		sum += score
	}

	avgScore := sum / float64(len(scores))

	fitness := 237.0 - avgScore

	return fitness, nil
}

// Mutate applies random changes to the weights and biases of the neural network
func (n *NeuralNetPlayer) Mutate(rng *rand.Rand) {
	mutateNode := func(node *gorgonia.Node) error {
		if node == nil || node.Value() == nil {
			return errors.New("node is nil or has no value")
		}

		data, ok := node.Value().Data().([]float64)
		if !ok {
			return errors.New("node value is not of type []float64")
		}

		for i := range data {
			if rng.Float64() < 0.02 {
				data[i] += rng.NormFloat64() * 0.1 // Small Gaussian perturbation
			}
			
		}
		return nil
	}

	// Mutate all weight and bias nodes
	_ = mutateNode(n.Weights1)
	_ = mutateNode(n.Biases1)
	_ = mutateNode(n.Weights2)
	_ = mutateNode(n.Biases2)
	_ = mutateNode(n.Weights3)
	_ = mutateNode(n.Biases3)
}

func (n *NeuralNetPlayer) Crossover(genome eaopt.Genome, rng *rand.Rand) {
	partner, ok := genome.(*NeuralNetPlayer)
	if !ok {
		return
	}

	crossoverNode := func(node1, node2 *gorgonia.Node) error {
		if node1 == nil || node2 == nil || node1.Value() == nil || node2.Value() == nil {
			return errors.New("one or both nodes are nil or have no value")
		}

		data1, ok1 := node1.Value().Data().([]float64)
		data2, ok2 := node2.Value().Data().([]float64)
		if !ok1 || !ok2 {
			return errors.New("node values are not of type []float64")
		}

		for i := range data1 {
			if rng.Float64() < 0.5 {
				data1[i], data2[i] = data2[i], data1[i] // Swap with 50% probability
			}
		}
		return nil
	}

	_ = crossoverNode(n.Weights1, partner.Weights1)
	_ = crossoverNode(n.Biases1, partner.Biases1)
	_ = crossoverNode(n.Weights2, partner.Weights2)
	_ = crossoverNode(n.Biases2, partner.Biases2)
	_ = crossoverNode(n.Weights3, partner.Weights3)
	_ = crossoverNode(n.Biases3, partner.Biases3)
}

func (n *NeuralNetPlayer) Clone() eaopt.Genome {
	return DeepCopyNeuralNetPlayer(n)
}
