package nnet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorgonia.org/gorgonia"
)

// Helper function to compare shapes and values
func compareNodes(t *testing.T, original, copy *gorgonia.Node) {
	t.Helper()

	// checking that nodes are non-nil
	require.NotNil(t, original, "Original node should not be nil")
	require.NotNil(t, copy, "Copied node should not be nil")

	// checking that shapes match
	assert.Equal(t, original.Shape(), copy.Shape(), "Shapes should be identical")

	// checking that values match
	var originalVal, copyVal gorgonia.Value
	gorgonia.Read(original, &originalVal)
	gorgonia.Read(copy, &copyVal)
	assert.Equal(t, originalVal, copyVal, "Copied node values should match original")
}

func TestDeepCopyNeuralNetPlayer(t *testing.T) {
	t.Parallel()

	original := NewNeuralNetPlayer(32, 32)
	copy  := DeepCopyNeuralNetPlayer(original)

	// make sure method executes successfully
	require.NotNil(t, copy, "Copied NeuralNetPlayer should not be nil")

	// make sure a deep copy was created
	assert.NotEqual(t, original.Graph, copy.Graph, "Graphs should be different instances")

	// compare nodes
	compareNodes(t, original.Input, copy.Input)
	compareNodes(t, original.Hidden1, copy.Hidden1)
	compareNodes(t, original.Hidden2, copy.Hidden2)
	compareNodes(t, original.Output, copy.Output)

	compareNodes(t, original.Weights1, copy.Weights1)
	compareNodes(t, original.Biases1, copy.Biases1)
	compareNodes(t, original.Weights2, copy.Weights2)
	compareNodes(t, original.Biases2, copy.Biases2)
	compareNodes(t, original.Weights3, copy.Weights3)
	compareNodes(t, original.Biases3, copy.Biases3)
}
