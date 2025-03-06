package game

import "fmt"

// A Category is one of the 9 ways you can score in Jumbleberry fields.
// It should take in a slice of Berries and calculate/store a score.
// It should also be able to return a score that has already been calculated.
type Category interface {
	CalcScore(berries []Berry) (int, error)
	GetScore() int
}

// A BaseCategory implements the basic functionality for a Category
type BaseCategory struct {
	Score int
	Used  bool
}

// GetScore returns the score for all
func (b BaseCategory) GetScore() int {
	return b.Score
}

func (b BaseCategory) String() string {
	if !b.Used {
		return "NOT USED, SCORE 0"
	}

	return fmt.Sprintf("USED, SCORE %d", b.Score)
}

// Scoring for the Jumbleberry Section is based on the amount of Jumbleberries rolled.
// Each Jumbleberry roll is worth 2 points.
// This means the maximum achievable score for this field is 10 points.
type JumbleberryCategory struct {
	BaseCategory
}

// CalcScore returns 2 * the number of Jumbleberries in the provided slice
func (j *JumbleberryCategory) CalcScore(berries []Berry) (int, error) {
	if len(berries) != 5 {
		return -1, fmt.Errorf("len of berries must be 5, got %d", len(berries))
	}

	if j.Used {
		return -1, fmt.Errorf("category has already been scored")
	}

	score := 0
	for _, berry := range berries {
		if berry == Jumbleberry {
			score += 2
		}
	}

	j.Score = score
	j.Used = true

	return score, nil
}

// Scoring for the Sugarberry Section is based on the amount of Sugarberries rolled.
// Each Sugarberry rolled is worth 2 points.
// This means the maximum achievable score for this field is 10 points.
type SugarberryCategory struct {
	BaseCategory
}

// CalcScore returns 2 * the number of Sugarberries in the provided slice
func (s *SugarberryCategory) CalcScore(berries []Berry) (int, error) {
	if len(berries) != 5 {
		return -1, fmt.Errorf("len of berries must be 5, got %d", len(berries))
	}

	if s.Used {
		return -1, fmt.Errorf("category has already been scored")
	}

	score := 0
	for _, berry := range berries {
		if berry == Sugarberry {
			score += 2
		}
	}

	s.Score = score
	s.Used = true

	return score, nil
}

// Scoring for the Pickleberry Section is based on the amount of Pickleberries rolled.
// Each Pickleberry rolled is worth 4 points.
// This means the maximum achievable score for this field is 20 points.
type PickleberryCategory struct {
	BaseCategory
}

// CalcScore returns 4 * the number of Pickleberries in the provided slice
func (p *PickleberryCategory) CalcScore(berries []Berry) (int, error) {
	if len(berries) != 5 {
		return -1, fmt.Errorf("len of berries must be 5, got %d", len(berries))
	}

	if p.Used {
		return -1, fmt.Errorf("category has already been scored")
	}

	score := 0
	for _, berry := range berries {
		if berry == Pickleberry {
			score += 4
		}
	}

	p.Score = score
	p.Used = true

	return score, nil
}

// Scoring for the Moonberry Section is based on the amount of Moonberries rolled.
// Each Moonberry rolled is worth 7 points.
// This means the maximum achievable score for this field is 35 points.
type MoonberryCategory struct {
	BaseCategory
}

// CalcScore returns 7 * the number of Moonberries in the provided slice
func (m *MoonberryCategory) CalcScore(berries []Berry) (int, error) {
	if len(berries) != 5 {
		return -1, fmt.Errorf("len of berries must be 5, got %d", len(berries))
	}

	if m.Used {
		return -1, fmt.Errorf("category has already been scored")
	}

	score := 0
	for _, berry := range berries {
		if berry == Moonberry {
			score += 7
		}
	}

	m.Score = score
	m.Used = true

	return score, nil
}

// In order to be able to score in this section you need to have at least 3 of one type of berry, or 3 pests.
// You can have more than 3 of one type of berry to score in this section.
// Scoring is based on the total points of all 5 dice.
// If you use three or more pests, the pests count as zero points and the only points received are those of the berries.
type ThreeCategory struct {
	BaseCategory
}

// CalcScore will return the combined score of all 5 berries/pests, as long as there are at least three of one type.
// If there are not three of one type, it will return 0.
func (t *ThreeCategory) CalcScore(berries []Berry) (int, error) {
	if len(berries) != 5 {
		return -1, fmt.Errorf("len of berries must be 5, got %d", len(berries))
	}

	if t.Used {
		return -1, fmt.Errorf("category has already been scored")
	}

	jberryCount, sberryCount, pberryCount, mberryCount, pestCount := 0, 0, 0, 0, 0
	score := 0

	for _, berry := range berries {
		switch berry {
		case Jumbleberry:
			score += 2
			jberryCount += 1
		case Sugarberry:
			score += 2
			sberryCount += 1
		case Pickleberry:
			score += 4
			pberryCount += 1
		case Moonberry:
			score += 7
			mberryCount += 1
		case Pest:
			pestCount += 1
		}
	}

	t.Used = true
	t.Score  = 0

	if jberryCount >= 3 || sberryCount >= 3 || pberryCount >= 3 || mberryCount >= 3 || pestCount >= 3 {
		t.Score = score
		return score, nil
	}

	return 0, nil
}

// In order to be able to score in this section you need to have at least 4 of one type of berry, or 4 pests.
// You can have more than 4 of one type of berry to score in this section.
// Scoring is based on the total points of all 5 dice.
// If you use four or more pests, the pests count as zero points and the only points received are those of the berries.
type FourCategory struct {
	BaseCategory
}

// CalcScore will return the combined score of all 5 berries/pests, as long as there are at least four of one type.
// If there are not four of one type, it will return 0.
func (f *FourCategory) CalcScore(berries []Berry) (int, error) {
	if len(berries) != 5 {
		return -1, fmt.Errorf("len of berries must be 5, got %d", len(berries))
	}

	if f.Used {
		return -1, fmt.Errorf("category has already been scored")
	}

	jberryCount, sberryCount, pberryCount, mberryCount, pestCount := 0, 0, 0, 0, 0
	score := 0

	for _, berry := range berries {
		switch berry {
		case Jumbleberry:
			score += 2
			jberryCount += 1
		case Sugarberry:
			score += 2
			sberryCount += 1
		case Pickleberry:
			score += 4
			pberryCount += 1
		case Moonberry:
			score += 7
			mberryCount += 1
		case Pest:
			pestCount += 1
		}
	}

	f.Used = true
	f.Score = 0

	if jberryCount >= 4 || sberryCount >= 4 || pberryCount >= 4 || mberryCount >= 4 || pestCount >= 4 {
		f.Score = score
		return score, nil
	}

	return 0, nil
}

// In order to be able to score in this section you need to have 5 berries of the same type.
// This field is the hardest one to score in and in many games is recorded as a zero.
// The easiest way to score this field is to try to get five Jumbleberries or five Sugarberries.
// This will give you a score of 10 for this field.
type FiveCategory struct {
	BaseCategory
}

// CalcScore will return 5 * berry value if there are five of the same type, or 0 if there aren't 5 of a kind.
func (f *FiveCategory) CalcScore(berries []Berry) (int, error) {
	if len(berries) != 5 {
		return -1, fmt.Errorf("len of berries must be 5, got %d", len(berries))
	}

	if f.Used {
		return -1, fmt.Errorf("category has already been scored")
	}

	f.Used = true

	// checking if all 5 berries are the same type
	// if they aren't return 0
	berryType := berries[0]
	for _, berry := range berries {
		if berry != berryType {
			f.Score = 0
			return 0, nil
		}
	}

	// if all 5 are the same type, determine how many points to score
	switch berryType {
	case Jumbleberry:
		f.Score = 10
		return 10, nil
	case Sugarberry:
		f.Score = 10
		return 10, nil
	case Pickleberry:
		f.Score = 20
		return 20, nil
	case Moonberry:
		f.Score = 35
		return 35, nil
	default:
		f.Score = 0
		return 0, nil
	}
}

// In order to be able to score in this section you must have one Jumbleberry, one Sugarberry, one Pickleberry, and one Moonberry.
// Scoring is based on the total of all 5 dice for this section so the maximum score for this section is 22 (if you have an extra Moonberry).
type MixedCategory struct {
	BaseCategory
}

// CalcScore returns the sum of the score of all berries if all four types of berries are present.
// If not all four types are present, it returns 0.
func (m *MixedCategory) CalcScore(berries []Berry) (int, error) {
	if len(berries) != 5 {
		return -1, fmt.Errorf("len of berries must be 5, got %d", len(berries))
	}

	if m.Used {
		return -1, fmt.Errorf("category has already been scored")
	}

	foundJB, foundSB, foundPB, foundMB := false, false, false, false
	score := 0

	for _, berry := range berries {
		switch berry {
		case Jumbleberry:
			foundJB = true
			score += 2
		case Sugarberry:
			foundSB = true
			score += 2
		case Pickleberry:
			foundPB = true
			score += 4
		case Moonberry:
			foundMB = true
			score += 7
		}
	}

	m.Used = true

	if !(foundJB && foundSB && foundPB && foundMB) {
		m.Score = 0
		return 0, nil
	}

	m.Score = score

	return score, nil
}

// This category adds up the score of all provided berries, regardless of what berries are present.
type FreeCategory struct {
	BaseCategory
}

// CalcScore returns the score of all of the berries added together, regardless fo the type
func (f *FreeCategory) CalcScore(berries []Berry) (int, error) {
	if len(berries) != 5 {
		return -1, fmt.Errorf("len of berries must be 5, got %d", len(berries))
	}

	if f.Used {
		return -1, fmt.Errorf("category has already been scored")
	}

	score := 0
	for _, berry := range berries {
		switch berry {
		case Jumbleberry:
			score += 2
		case Sugarberry:
			score += 2
		case Pickleberry:
			score += 4
		case Moonberry:
			score += 7
		}
	}

	f.Used = true
	f.Score = score

	return score, nil
}
