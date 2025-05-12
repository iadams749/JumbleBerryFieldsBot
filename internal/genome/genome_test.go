package genome

import (
	"fmt"
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		genome *Genome
	}{
		{
			name: "Simple case",
			genome: &Genome{
				Biases:           [][]float64{{0.1, 0.2}, {0.3, 0.4}},
				HiddenLayerSizes: []int{2, 3},
				Weights:          [][][]float64{{{0.5, 0.6}, {0.7, 0.8}}, {{0.9, 1.0}, {1.1, 1.2}}},
			},
		},
		{
			name: "Larger case",
			genome: &Genome{
				Biases:           [][]float64{{0.5, 0.6, 0.7}, {0.8, 0.9, 1.0}},
				HiddenLayerSizes: []int{4, 5, 6},
				Weights:          [][][]float64{{{1.1, 1.2}, {1.3, 1.4}}, {{1.5, 1.6}, {1.7, 1.8}}, {{1.1, 1.2}, {1.3, 1.4}}, {{1.5, 1.6}, {1.7, 1.8}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			copy := tt.genome.Clone()

			copyG, ok := copy.(*Genome)
			if !ok {
				t.Errorf("Failed to cast as Genome")
			}

			if !reflect.DeepEqual(tt.genome, copyG) {
				t.Errorf("Clone failed: copied Genome is not identical to the original")
				fmt.Println(tt.genome)
				fmt.Println(copyG)
			}

			if len(tt.genome.Biases) > 0 {
				copyG.Biases[0][0] = 99.99
				if tt.genome.Biases[0][0] == 99.99 {
					t.Errorf("Clone failed: Biases were not deeply copied")
				}
			}

			if len(tt.genome.HiddenLayerSizes) > 0 {
				copyG.HiddenLayerSizes[0] = 99
				if tt.genome.HiddenLayerSizes[0] == 99 {
					t.Errorf("Clone failed: HiddenLayerSizes were not deeply copied")
				}
			}

			if len(tt.genome.Weights) > 0 && len(tt.genome.Weights[0]) > 0 && len(tt.genome.Weights[0][0]) > 0 {
				copyG.Weights[0][0][0] = 99.99
				if tt.genome.Weights[0][0][0] == 99.99 {
					t.Errorf("Clone failed: Weights were not deeply copied")
				}
			}
		})
	}
}
