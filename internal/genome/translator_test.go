package genome

import (
	"reflect"
	"testing"

	"github.com/iadams749/JumbleBerryFieldsBot/internal/game"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func TestGetTopKValues(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		data     []float64
		k        int
		expected []float64
		indices  []int
		wantErr  bool
	}{
		{
			name:     "Normal case",
			data:     []float64{1.0, 3.0, 2.0, 5.0, 4.0},
			k:        3,
			expected: []float64{5.0, 4.0, 3.0},
			indices:  []int{3, 4, 1},
			wantErr:  false,
		},
		{
			name:     "K greater than length",
			data:     []float64{10.0, 20.0},
			k:        5,
			expected: []float64{20.0, 10.0},
			indices:  []int{1, 0},
			wantErr:  false,
		},
		{
			name:     "K equals 0",
			data:     []float64{1.0, 2.0, 3.0},
			k:        0,
			expected: []float64{},
			indices:  []int{},
			wantErr:  false,
		},
		{
			name:     "Not a Tensor",
			data:     nil,
			expected: nil,
			indices:  nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var node *gorgonia.Node

			if tt.data != nil {
				tensorData := tensor.New(tensor.WithBacking(tt.data), tensor.WithShape(len(tt.data)))
				node = gorgonia.NewTensor(gorgonia.NewGraph(), tensor.Float64, 1, gorgonia.WithValue(tensorData))
			} else {
				node = &gorgonia.Node{}
			}

			values, indices, err := GetTopKValues(node, tt.k)

			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}

			if !equalSlices(values, tt.expected) || !equalSlicesInt(indices, tt.indices) {
				t.Errorf("expected values %v with indices %v, got values %v with indices %v", tt.expected, tt.indices, values, indices)
			}
		})
	}
}

// equalSlices checks if two float64 slices have the same length and elements.
func equalSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// equalSlicesInt checks if two integer slices have the same length and elements.
func equalSlicesInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTranslateGameState(t *testing.T) {
	t.Parallel()
	type args struct {
		gs *game.GameState
	}
	tests := []struct {
		name string
		args args
		want *tensor.Dense
	}{
		{
			name: "All Jumbleberries/No Categories Used",
			args: args{
				gs: &game.GameState{
					Jars: []*game.Jar{
						{
							Berry:  game.Jumbleberry,
							Rolled: true,
						},
						{
							Berry:  game.Jumbleberry,
							Rolled: true,
						},
						{
							Berry:  game.Jumbleberry,
							Rolled: true,
						},
						{
							Berry:  game.Jumbleberry,
							Rolled: true,
						},
						{
							Berry:  game.Jumbleberry,
							Rolled: true,
						},
					},
					RollsLeftInTurn: 2,
					Categories:      game.NewGame().Categories,
				},
			},
			want: tensor.New(
				tensor.WithBacking([]float64{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}),
				tensor.WithShape(1, 37), // Shape matches the input layer
			),
		},
		{
			name: "All Berry Types/No Categories Used",
			args: args{
				gs: &game.GameState{
					Jars: []*game.Jar{
						{
							Berry:  game.Jumbleberry,
							Rolled: true,
						},
						{
							Berry:  game.Sugarberry,
							Rolled: true,
						},
						{
							Berry:  game.Pickleberry,
							Rolled: true,
						},
						{
							Berry:  game.Moonberry,
							Rolled: true,
						},
						{
							Berry:  game.Pest,
							Rolled: true,
						},
					},
					RollsLeftInTurn: 1,
					Categories:      game.NewGame().Categories,
				},
			},
			want: tensor.New(
				tensor.WithBacking([]float64{1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}),
				tensor.WithShape(1, 37), // Shape matches the input layer
			),
		},
		{
			name: "All Berry Types/All Categories Used",
			args: args{
				gs: func() *game.GameState {
					state := &game.GameState{
						Jars: []*game.Jar{
							{
								Berry:  game.Jumbleberry,
								Rolled: true,
							},
							{
								Berry:  game.Sugarberry,
								Rolled: true,
							},
							{
								Berry:  game.Pickleberry,
								Rolled: true,
							},
							{
								Berry:  game.Moonberry,
								Rolled: true,
							},
							{
								Berry:  game.Pest,
								Rolled: true,
							},
						},
						RollsLeftInTurn: 0,
						Categories:      game.NewGame().Categories,
					}

					state.Categories.JumbleberryCategory.Used = true
					state.Categories.SugarberryCategory.Used = true
					state.Categories.PickleberryCategory.Used = true
					state.Categories.MoonberryCategory.Used = true
					state.Categories.ThreeCategory.Used = true
					state.Categories.FourCategory.Used = true
					state.Categories.FiveCategory.Used = true
					state.Categories.MixedCategory.Used = true
					state.Categories.FreeCategory.Used = true

					return state
				}(),
			},
			want: tensor.New(
				tensor.WithBacking([]float64{1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}),
				tensor.WithShape(1, 37), // Shape matches the input layer
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := TranslateGameState(tt.args.gs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TranslateGameState() = %v, want %v", got, tt.want)
			}
		})
	}
}

// This unit test makes sure that the function panics when it is expected to panic
func TestTranslateGameState_Panic(t *testing.T) {
	t.Parallel()
	type args struct {
		gs *game.GameState
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Incorrect Number of Jars",
			args: args{
				gs: &game.GameState{},
			},
		},
		{
			name: "Invalid Berry",
			args: args{
				gs: &game.GameState{
					Jars: []*game.Jar{
						{
							Berry:  game.Jumbleberry,
							Rolled: true,
						},
						{
							Berry:  game.Sugarberry,
							Rolled: true,
						},
						{
							Berry:  game.Pickleberry,
							Rolled: true,
						},
						{
							Berry:  game.Moonberry,
							Rolled: true,
						},
						{
							Berry:  -1,
							Rolled: true,
						},
					},
				},
			},
		},
		{
			name: "Jars Haven't Been Rolled Yet",
			args: args{
				gs: game.NewGame(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("expected panic, but function did not panic")
				}
			}()

			TranslateGameState(tt.args.gs)
		})
	}
}
