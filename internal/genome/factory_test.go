package genome

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestNewGenome(t *testing.T) {
	t.Parallel()
	type args struct {
		rng              *rand.Rand
		hiddenLayerSizes []int
	}
	tests := []struct {
		name     string
		args     args
		fileName string
	}{
		{
			name: "32 x 32 Hidden Layers",
			args: args{
				rng: rand.New(rand.NewSource(0)),
				hiddenLayerSizes: []int{32,32},
			},
			fileName: "testdata/32_32.json",
		},
		{
			name: "2 x 4 x 8 Hidden Layers",
			args: args{
				rng: rand.New(rand.NewSource(0)),
				hiddenLayerSizes: []int{2,4,8},
			},
			fileName: "testdata/2_4_8.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Open the JSON file
			file, err := os.Open(tt.fileName)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()

			// Read the file contents
			bytes, err := io.ReadAll(file)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			// Unmarshal JSON into struct
			var want Genome
			err = json.Unmarshal(bytes, &want)
			if err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}

			if got := NewGenome(tt.args.rng, tt.args.hiddenLayerSizes); !reflect.DeepEqual(got, &want) {
				t.Errorf("NewGenome() = %v, want %v", got, want)
			}
		})
	}
}

func TestNewGenomeFactory(t *testing.T) {
	t.Parallel()
	type args struct {
		rng              *rand.Rand
		hiddenLayerSizes []int
	}
	tests := []struct {
		name     string
		args     args
		fileName string
	}{
		{
			name: "32 x 32 Hidden Layers",
			args: args{
				rng: rand.New(rand.NewSource(0)),
				hiddenLayerSizes: []int{32,32},
			},
			fileName: "testdata/32_32.json",
		},
		{
			name: "2 x 4 x 8 Hidden Layers",
			args: args{
				rng: rand.New(rand.NewSource(0)),
				hiddenLayerSizes: []int{2,4,8},
			},
			fileName: "testdata/2_4_8.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Open the JSON file
			file, err := os.Open(tt.fileName)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()

			// Read the file contents
			bytes, err := io.ReadAll(file)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			// Unmarshal JSON into struct
			var want Genome
			err = json.Unmarshal(bytes, &want)
			if err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}

			factory := NewGenomeFactory(tt.args.hiddenLayerSizes)

			if got := factory(tt.args.rng); !reflect.DeepEqual(got, &want) {
				t.Errorf("NewGenomeFactory() = %v, want %v", got, want)
			}
		})
	}
}