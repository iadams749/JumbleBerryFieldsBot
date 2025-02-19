package game

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want *GameState
	}{
		{
			name: "New Game",
			want: &GameState{
				Jars:            []*Jar{{}, {}, {}, {}, {}},
				RollsLeftInTurn: 3,
				Categories: GameCategories{
					JumbleberryCategory: &JumbleberryCategory{},
					SugarberryCategory:  &SugarberryCategory{},
					PickleberryCategory: &PickleberryCategory{},
					MoonberryCategory:   &MoonberryCategory{},
					ThreeCategory:       &ThreeCategory{},
					FourCategory:        &FourCategory{},
					FiveCategory:        &FiveCategory{},
					MixedCategory:       &MixedCategory{},
					FreeCategory:        &FreeCategory{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewGame(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameState_RollJars_Success(t *testing.T) {
	t.Parallel()
	gs := NewGame()

	err := gs.RollJars()

	if err != nil {
		t.Errorf("expected nil error but for %e", err)
	}

	for _, jar := range gs.Jars {
		if !jar.Rolled {
			t.Errorf("jar is not rolled but should be")
		}
	}

	if gs.RollsLeftInTurn != 2 {
		t.Errorf("rolls left wasn't decreased")
	}
}

func TestGameState_RollJars_NoRollLeft(t *testing.T) {
	t.Parallel()
	gs := NewGame()

	gs.RollsLeftInTurn = 0

	err := gs.RollJars()

	if err == nil {
		t.Errorf("expect error but got nil")
	}
}

func TestGameState_LockJar(t *testing.T) {
	t.Parallel()
	gs := NewGame()

	gs.LockJar(0)

	if !gs.Jars[0].Locked {
		t.Errorf("jar is not locked but should be")
	}
}

func TestGameState_String(t *testing.T) {
	gs := NewGame()

	got := gs.String()

	want := "ROUND: 1\nROLLS LEFT: 3\nSCORE: 0\nJars: [EMPTY EMPTY EMPTY EMPTY EMPTY]\nJumbleberry: NOT USED, SCORE 0\nSugarberry: NOT USED, SCORE 0\nPickleberry: NOT USED, SCORE 0\nMoonberry: NOT USED, SCORE 0\nThree of a Kind: NOT USED, SCORE 0\nFour of a Kind: NOT USED, SCORE 0\nFive of a Kind: NOT USED, SCORE 0\nMixed Basket: NOT USED, SCORE 0\nFree Roll: NOT USED, SCORE 0\n"
	if got != want {
		t.Errorf("incorrect string, got: %s", got)
	}
}
