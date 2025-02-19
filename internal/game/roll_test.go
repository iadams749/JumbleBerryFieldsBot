package game

import (
	"testing"
)

// this test will do 10000 random rolls and ensure that the number of each result is within
// 5 standard deviations of what is expected
// there is a VERY small chance this test will fail by random chance
func Test_doRoll(t *testing.T) {
	t.Parallel()

	berryCounts := make(map[Berry]int)

	for range 10000 {
		result := doRoll()
		if count, ok := berryCounts[result]; ok {
			berryCounts[result] = count + 1
		} else {
			berryCounts[result] = 1
		}
	}

	if berryCounts[Pest] < 850 || berryCounts[Pest] > 1150 {
		t.Errorf("Pest count is outside expected range of 850-1150: %d", berryCounts[Pest])
	}
	if berryCounts[Moonberry] < 850 || berryCounts[Moonberry] > 1150 {
		t.Errorf("Moonberry count is outside expected range of 850-1150: %d", berryCounts[Moonberry])
	}
	if berryCounts[Pickleberry] < 1800 || berryCounts[Pickleberry] > 2200 {
		t.Errorf("Pickleberry count is outside expected range of 1800-2200: %d", berryCounts[Pickleberry])
	}
	if berryCounts[Jumbleberry] < 2750 || berryCounts[Jumbleberry] > 3250 {
		t.Errorf("Jumbleberry count is outside expected range of 2750-3250: %d", berryCounts[Jumbleberry])
	}
	if berryCounts[Sugarberry] < 2750 || berryCounts[Sugarberry] > 3250 {
		t.Errorf("Sugarberry count is outside expected range of 2750-3250: %d", berryCounts[Sugarberry])
	}
}

func TestDoRolls(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "5 rolls",
			args: args{
				n: 5,
			},
		},
		{
			name: "3 rolls",
			args: args{
				n: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DoRolls(tt.args.n)

			if len(got) != tt.args.n {
				t.Errorf("expected to get slice of len %d, got %d", tt.args.n, len(got))
			}

			for _, berry := range got {
				if berry < 0 || berry > 4 {
					t.Errorf("invalid berry value: %d", berry)
				}
			}
		})
	}
}
