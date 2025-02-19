package game

import "fmt"

// A GameState represents the state for a single game.
// It is comprised of the score, jars, rounds left, rolls left and categories.
type GameState struct {
	Categories      GameCategories
	Jars            []*Jar
	RollsLeftInTurn int
	RoundsCompleted int
	Score           int
}

// A GameCategories object represents the 9 different scoring categories.
type GameCategories struct {
	JumbleberryCategory *JumbleberryCategory
	SugarberryCategory  *SugarberryCategory
	PickleberryCategory *PickleberryCategory
	MoonberryCategory   *MoonberryCategory
	ThreeCategory       *ThreeCategory
	FourCategory        *FourCategory
	FiveCategory        *FiveCategory
	MixedCategory       *MixedCategory
	FreeCategory        *FreeCategory
}

// NewGame returns a *GameState in the starting state
func NewGame() *GameState {
	return &GameState{
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
	}
}

// RollJars rolls all jars in a game state.
// If a Jar is locked, then it won't be rolled.
func (gs *GameState) RollJars() error {
	if gs.RollsLeftInTurn < 1 {
		return fmt.Errorf("no rolls left")
	}

	for _, jar := range gs.Jars {
		jar.Roll()
	}

	gs.RollsLeftInTurn -= 1
	return nil
}

func (gs *GameState) GetBerries() []Berry {
	var berries []Berry

	for _, jar := range(gs.Jars) {
		berries = append(berries, jar.Berry)
	}

	return berries
}

func (gs *GameState) NewTurn() {
	for _, jar := range(gs.Jars) {
		jar.Reset()
	}

	gs.RollsLeftInTurn = 3
	gs.RoundsCompleted += 1
}

// LockJar will lock a given jar in a GameState.
// If a Jar is already locked, it will stay locked
func (gs *GameState) LockJar(n int) {
	gs.Jars[n].Lock()
}

// UnlockJar will unlock a given jar in a GameState.
// If a Jar is already unlocked, it will stay unlocked
func (gs *GameState) UnlockJar(n int) {
	gs.Jars[n].Unlock()
}

func (gs *GameState) String() string {
	str := fmt.Sprintf("ROUND: %d\n", gs.RoundsCompleted+1)
	str += fmt.Sprintf("ROLLS LEFT: %d\n", gs.RollsLeftInTurn)
	str += fmt.Sprintf("SCORE: %d\n", gs.Score)
	str += "Jars: " + fmt.Sprint(gs.Jars) + "\n"
	str += "Jumbleberry: " + gs.Categories.JumbleberryCategory.String() + "\n"
	str += "Sugarberry: " + gs.Categories.SugarberryCategory.String() + "\n"
	str += "Pickleberry: " + gs.Categories.PickleberryCategory.String() + "\n"
	str += "Moonberry: " + gs.Categories.MoonberryCategory.String() + "\n"
	str += "Three of a Kind: " + gs.Categories.ThreeCategory.String() + "\n"
	str += "Four of a Kind: " + gs.Categories.FourCategory.String() + "\n"
	str += "Five of a Kind: " + gs.Categories.FiveCategory.String() + "\n"
	str += "Mixed Basket: " + gs.Categories.MixedCategory.String() + "\n"
	str += "Free Roll: " + gs.Categories.FreeCategory.String() + "\n"
	return str
}
