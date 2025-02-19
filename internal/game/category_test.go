package game

import (
	"testing"
)

func TestBaseCategory_GetScore(t *testing.T) {
	type fields struct {
		Score int
		Used  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "New BaseCategory",
			fields: fields{
				Score: 0,
				Used:  false,
			},
			want: 0,
		},
		{
			name: "Completed BaseCategory",
			fields: fields{
				Score: 10,
				Used:  true,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseCategory{
				Score: tt.fields.Score,
				Used:  tt.fields.Used,
			}
			if got := b.GetScore(); got != tt.want {
				t.Errorf("BaseCategory.GetScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseCategory_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		Score int
		Used  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Unused",
			fields: fields{
				Used: false,
			},
			want: "NOT USED, SCORE 0",
		},
		{
			name: "Used - Score 5",
			fields: fields{
				Used:  true,
				Score: 5,
			},
			want: "USED, SCORE 5",
		},
		{
			name: "Used - Score 10",
			fields: fields{
				Used:  true,
				Score: 10,
			},
			want: "USED, SCORE 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := BaseCategory{
				Score: tt.fields.Score,
				Used:  tt.fields.Used,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("BaseCategory.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJumbleberryCategory_CalcScore(t *testing.T) {
	t.Parallel()
	type fields struct {
		BaseCategory BaseCategory
	}
	type args struct {
		berries []Berry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Valid case - 3 Jumbleberries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Pest, Pickleberry},
			},
			want:    6, // 3 * 2 points
			wantErr: false,
		},
		{
			name: "Valid case - 5 Jumbleberries (max score)",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry},
			},
			want:    10, // 5 * 2 points
			wantErr: false,
		},
		{
			name: "Valid case - No Jumbleberries",
			args: args{
				berries: []Berry{Pest, Pest, Pest, Pest, Pest},
			},
			want:    0, // No points
			wantErr: false,
		},
		{
			name: "Invalid case - Too few berries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry},
			},
			want:    -1, // Should return error
			wantErr: true,
		},
		{
			name: "Invalid case - Too many berries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry, Pest},
			},
			want:    -1, // Should return error
			wantErr: true,
		},
		{
			name: "Invalid case - Already Used",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry, Sugarberry},
			},
			fields: fields{
				BaseCategory: BaseCategory{
					Used: true,
				},
			},
			want:    -1, // Should return error due being already used
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			j := &JumbleberryCategory{
				BaseCategory: tt.fields.BaseCategory,
			}
			got, err := j.CalcScore(tt.args.berries)
			if (err != nil) != tt.wantErr {
				t.Errorf("JumbleberryCategory.CalcScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JumbleberryCategory.CalcScore() = %v, want %v", got, tt.want)
			}
			if err == nil && j.Used == false {
				t.Errorf("category should be marked used but is not")
			}
		})
	}
}

func TestSugarberryCategory_CalcScore(t *testing.T) {
	t.Parallel()
	type fields struct {
		BaseCategory BaseCategory
	}
	type args struct {
		berries []Berry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Valid case - 3 Sugarberries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry, Sugarberry, Pest, Pickleberry},
			},
			want:    6, // 3 * 2 points
			wantErr: false,
		},
		{
			name: "Valid case - 5 Sugarberries (max score)",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry, Sugarberry, Sugarberry, Sugarberry},
			},
			want:    10, // 5 * 2 points
			wantErr: false,
		},
		{
			name: "Valid case - No Sugarberries",
			args: args{
				berries: []Berry{Pest, Pest, Jumbleberry, Pest, Moonberry},
			},
			want:    0, // No points
			wantErr: false,
		},
		{
			name: "Invalid case - Too few berries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry},
			},
			want:    -1, // Should return error
			wantErr: true,
		},
		{
			name: "Invalid case - Too many berries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry, Sugarberry, Sugarberry, Sugarberry, Pest},
			},
			want:    -1, // Should return error
			wantErr: true,
		},
		{
			name: "Invalid case - Already Used",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry, Sugarberry},
			},
			fields: fields{
				BaseCategory: BaseCategory{
					Used: true,
				},
			},
			want:    -1, // Should return error due being already used
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &SugarberryCategory{
				BaseCategory: tt.fields.BaseCategory,
			}
			got, err := s.CalcScore(tt.args.berries)
			if (err != nil) != tt.wantErr {
				t.Errorf("SugarberryCategory.CalcScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SugarberryCategory.CalcScore() = %v, want %v", got, tt.want)
			}
			if err == nil && s.Used == false {
				t.Errorf("category should be marked used but is not")
			}
		})
	}
}

func TestPickleberryCategory_CalcScore(t *testing.T) {
	t.Parallel()
	type fields struct {
		BaseCategory BaseCategory
	}
	type args struct {
		berries []Berry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Valid case - 3 Pickleberries",
			args: args{
				berries: []Berry{Pickleberry, Pickleberry, Sugarberry, Pest, Pickleberry},
			},
			want:    12, // 3 * 4 points
			wantErr: false,
		},
		{
			name: "Valid case - 5 Pickleberries (max score)",
			args: args{
				berries: []Berry{Pickleberry, Pickleberry, Pickleberry, Pickleberry, Pickleberry},
			},
			want:    20, // 5 * 4 points
			wantErr: false,
		},
		{
			name: "Valid case - No Pickleberries",
			args: args{
				berries: []Berry{Pest, Pest, Jumbleberry, Pest, Moonberry},
			},
			want:    0, // No points
			wantErr: false,
		},
		{
			name: "Invalid case - Too few berries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry},
			},
			want:    -1, // Should return error
			wantErr: true,
		},
		{
			name: "Invalid case - Too many berries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry, Sugarberry, Sugarberry, Sugarberry, Pest},
			},
			want:    -1, // Should return error
			wantErr: true,
		},
		{
			name: "Invalid case - Already Used",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry, Sugarberry},
			},
			fields: fields{
				BaseCategory: BaseCategory{
					Used: true,
				},
			},
			want:    -1, // Should return error due being already used
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := &PickleberryCategory{
				BaseCategory: tt.fields.BaseCategory,
			}
			got, err := p.CalcScore(tt.args.berries)
			if (err != nil) != tt.wantErr {
				t.Errorf("PickleberryCategory.CalcScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PickleberryCategory.CalcScore() = %v, want %v", got, tt.want)
			}
			if err == nil && p.Used == false {
				t.Errorf("category should be marked used but is not")
			}
		})
	}
}

func TestMoonberryCategory_CalcScore(t *testing.T) {
	t.Parallel()
	type fields struct {
		BaseCategory BaseCategory
	}
	type args struct {
		berries []Berry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Valid case - 3 Moonberries",
			args: args{
				berries: []Berry{Moonberry, Moonberry, Sugarberry, Pest, Moonberry},
			},
			want:    21, // 3 * 7 points
			wantErr: false,
		},
		{
			name: "Valid case - 5 Moonberries (max score)",
			args: args{
				berries: []Berry{Moonberry, Moonberry, Moonberry, Moonberry, Moonberry},
			},
			want:    35, // 5 * 7 points
			wantErr: false,
		},
		{
			name: "Valid case - No Moonberries",
			args: args{
				berries: []Berry{Pest, Pest, Jumbleberry, Pest, Pickleberry},
			},
			want:    0, // No points
			wantErr: false,
		},
		{
			name: "Invalid case - Too few berries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry},
			},
			want:    -1, // Should return error
			wantErr: true,
		},
		{
			name: "Invalid case - Too many berries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry, Sugarberry, Sugarberry, Sugarberry, Pest},
			},
			want:    -1, // Should return error
			wantErr: true,
		},
		{
			name: "Invalid case - Already Used",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry, Sugarberry},
			},
			fields: fields{
				BaseCategory: BaseCategory{
					Used: true,
				},
			},
			want:    -1, // Should return error due being already used
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := &MoonberryCategory{
				BaseCategory: tt.fields.BaseCategory,
			}
			got, err := m.CalcScore(tt.args.berries)
			if (err != nil) != tt.wantErr {
				t.Errorf("MoonberryCategory.CalcScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MoonberryCategory.CalcScore() = %v, want %v", got, tt.want)
			}
			if err == nil && m.Used == false {
				t.Errorf("category should be marked used but is not")
			}
		})
	}
}

func TestThreeCategory_CalcScore(t *testing.T) {
	t.Parallel()
	type fields struct {
		BaseCategory BaseCategory
	}
	type args struct {
		berries []Berry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Valid case - 3 Jumbleberries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Pest, Pickleberry},
			},
			want:    10, // 3 Jumbleberries (2 points each) + 4 points from Pickleberry
			wantErr: false,
		},
		{
			name: "Valid case - 3 Sugarberries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry, Sugarberry, Pest, Moonberry},
			},
			want:    13, // 3 Sugarberries (2 points each) + 7 points from Moonberry
			wantErr: false,
		},
		{
			name: "Valid case - 3 Pest",
			args: args{
				berries: []Berry{Pest, Pest, Pest, Jumbleberry, Moonberry},
			},
			want:    9, // 7 points from Moonberry + 2 points from Jumbleberry (pests count as 0)
			wantErr: false,
		},
		{
			name: "Valid case - No 3 of a single type of berry",
			args: args{
				berries: []Berry{Jumbleberry, Pickleberry, Sugarberry, Moonberry, Pest},
			},
			want:    0, // No three of the same berry
			wantErr: false,
		},
		{
			name: "Invalid case - Not exactly 5 berries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Sugarberry, Pest},
			},
			want:    -1, // Should return error due to incorrect number of berries
			wantErr: true,
		},
		{
			name: "Invalid case - More than 5 berries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry, Sugarberry, Pest},
			},
			want:    -1, // Should return error due to incorrect number of berries
			wantErr: true,
		},
		{
			name: "Invalid case - Already Used",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry, Sugarberry},
			},
			fields: fields{
				BaseCategory: BaseCategory{
					Used: true,
				},
			},
			want:    -1, // Should return error due being already used
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tc := &ThreeCategory{
				BaseCategory: tt.fields.BaseCategory,
			}
			got, err := tc.CalcScore(tt.args.berries)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThreeCategory.CalcScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ThreeCategory.CalcScore() = %v, want %v", got, tt.want)
			}
			if err == nil && tc.Used == false {
				t.Errorf("category should be marked used but is not")
			}
		})
	}
}

func TestFourCategory_CalcScore(t *testing.T) {
	t.Parallel()
	type fields struct {
		BaseCategory BaseCategory
	}
	type args struct {
		berries []Berry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Valid case - 4 Jumbleberries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry, Pest},
			},
			want:    8, // 4 Jumbleberries (2 points each)
			wantErr: false,
		},
		{
			name: "Valid case - 4 Sugarberries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry, Sugarberry, Sugarberry, Pickleberry},
			},
			want:    12, // 4 Sugarberries (2 points each) + 4 points from Pickleberry
			wantErr: false,
		},
		{
			name: "Valid case - 4 Pests",
			args: args{
				berries: []Berry{Pest, Pest, Pest, Pest, Jumbleberry},
			},
			want:    2, // 2 points from Jumbleberry (pests count as 0)
			wantErr: false,
		},
		{
			name: "Valid case - More than 4 of the same berry",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry},
			},
			want:    10, // 5 Jumbleberries (2 points each)
			wantErr: false,
		},
		{
			name: "Valid case - No 4 of a single type of berry",
			args: args{
				berries: []Berry{Jumbleberry, Sugarberry, Pickleberry, Moonberry, Pest},
			},
			want:    0, // No four of the same berry
			wantErr: false,
		},
		{
			name: "Invalid case - Not exactly 5 berries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Sugarberry, Pest},
			},
			want:    -1, // Should return error due to incorrect number of berries
			wantErr: true,
		},
		{
			name: "Invalid case - More than 5 berries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry, Sugarberry, Pest},
			},
			want:    -1, // Should return error due to incorrect number of berries
			wantErr: true,
		},
		{
			name: "Invalid case - Already Used",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry, Sugarberry},
			},
			fields: fields{
				BaseCategory: BaseCategory{
					Used: true,
				},
			},
			want:    -1, // Should return error due being already used
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fc := &FourCategory{
				BaseCategory: tt.fields.BaseCategory,
			}
			got, err := fc.CalcScore(tt.args.berries)
			if (err != nil) != tt.wantErr {
				t.Errorf("FourCategory.CalcScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FourCategory.CalcScore() = %v, want %v", got, tt.want)
			}
			if err == nil && fc.Used == false {
				t.Errorf("category should be marked used but is not")
			}
		})
	}
}

func TestFiveCategory_CalcScore(t *testing.T) {
	t.Parallel()
	type fields struct {
		BaseCategory BaseCategory
	}
	type args struct {
		berries []Berry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Valid case - 5 Jumbleberries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry},
			},
			want:    10, // 5 Jumbleberries (10 points)
			wantErr: false,
		},
		{
			name: "Valid case - 5 Sugarberries",
			args: args{
				berries: []Berry{Sugarberry, Sugarberry, Sugarberry, Sugarberry, Sugarberry},
			},
			want:    10, // 5 Sugarberries (10 points)
			wantErr: false,
		},
		{
			name: "Valid case - 5 Pickleberries",
			args: args{
				berries: []Berry{Pickleberry, Pickleberry, Pickleberry, Pickleberry, Pickleberry},
			},
			want:    20, // 5 Pickleberries (20 points)
			wantErr: false,
		},
		{
			name: "Valid case - 5 Moonberries",
			args: args{
				berries: []Berry{Moonberry, Moonberry, Moonberry, Moonberry, Moonberry},
			},
			want:    35, // 5 Moonberries (35 points)
			wantErr: false,
		},
		{
			name: "Valid case - 5 Pests",
			args: args{
				berries: []Berry{Pest, Pest, Pest, Pest, Pest},
			},
			want:    0, // 5 Pests (0 points)
			wantErr: false,
		},
		{
			name: "Valid case - Mixed berries",
			args: args{
				berries: []Berry{Jumbleberry, Sugarberry, Pickleberry, Moonberry, Jumbleberry},
			},
			want:    0, // Not all the same berry
			wantErr: false,
		},
		{
			name: "Invalid case - Not exactly 5 berries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry},
			},
			want:    -1, // Should return error due to incorrect number of berries
			wantErr: true,
		},
		{
			name: "Invalid case - More than 5 berries",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry},
			},
			want:    -1, // Should return error due to incorrect number of berries
			wantErr: true,
		},
		{
			name: "Invalid case - Already Used",
			args: args{
				berries: []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Sugarberry, Sugarberry},
			},
			fields: fields{
				BaseCategory: BaseCategory{
					Used: true,
				},
			},
			want:    -1, // Should return error due being already used
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fc := &FiveCategory{
				BaseCategory: tt.fields.BaseCategory,
			}
			got, err := fc.CalcScore(tt.args.berries)
			if (err != nil) != tt.wantErr {
				t.Errorf("FiveCategory.CalcScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FiveCategory.CalcScore() = %v, want %v", got, tt.want)
			}
			if err == nil && fc.Used == false {
				t.Errorf("category should be marked used but is not")
			}
		})
	}
}

func TestMixedCategory_CalcScore(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		berries  []Berry
		used     bool
		expected int
		wantErr  bool
	}{
		{
			name:     "Valid case - All required berries",
			berries:  []Berry{Jumbleberry, Sugarberry, Pickleberry, Moonberry, Pest},
			expected: 15, // 2 + 2 + 4 + 7 = 15 points (all required berries)
			wantErr:  false,
		},
		{
			name:     "Valid case - All required berries with an extra Moonberry",
			berries:  []Berry{Jumbleberry, Sugarberry, Pickleberry, Moonberry, Moonberry},
			expected: 22, // 2 + 2 + 4 + 7 + 7 = 22 points (all required berries + extra Moonberry)
			wantErr:  false,
		},
		{
			name:     "Valid case - All required berries with an extra Jumbleberry",
			berries:  []Berry{Jumbleberry, Sugarberry, Pickleberry, Moonberry, Jumbleberry},
			expected: 17, // 2 + 2 + 4 + 7 + 2 = 17 points (all required berries + extra Jumbleberry)
			wantErr:  false,
		},
		{
			name:     "Invalid case - Missing Jumbleberry",
			berries:  []Berry{Sugarberry, Pickleberry, Moonberry, Moonberry, Pest},
			expected: 0, // Missing Jumbleberry
			wantErr:  false,
		},
		{
			name:     "Invalid case - Missing Sugarberry",
			berries:  []Berry{Jumbleberry, Pickleberry, Moonberry, Moonberry, Pest},
			expected: 0, // Missing Sugarberry
			wantErr:  false,
		},
		{
			name:     "Invalid case - Missing Pickleberry",
			berries:  []Berry{Jumbleberry, Sugarberry, Moonberry, Moonberry, Pest},
			expected: 0, // Missing Pickleberry
			wantErr:  false,
		},
		{
			name:     "Invalid case - Missing Moonberry",
			berries:  []Berry{Jumbleberry, Sugarberry, Pickleberry, Pest, Pest},
			expected: 0, // Missing Moonberry
			wantErr:  false,
		},
		{
			name:     "Invalid case - Not exactly 5 berries",
			berries:  []Berry{Jumbleberry, Sugarberry, Pickleberry},
			expected: -1, // Incorrect number of berries
			wantErr:  true,
		},
		{
			name:     "Invalid case - More than 5 berries",
			berries:  []Berry{Jumbleberry, Sugarberry, Pickleberry, Moonberry, Pest, Pickleberry},
			expected: -1, // Incorrect number of berries
			wantErr:  true,
		},
		{
			name:     "Invalid case - Already Used",
			berries:  []Berry{Sugarberry, Pickleberry, Moonberry, Moonberry, Pest},
			expected: -1,
			used:     true,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mc := &MixedCategory{
				BaseCategory: BaseCategory{
					Used: tt.used,
				},
			}
			got, err := mc.CalcScore(tt.berries)
			if (err != nil) != tt.wantErr {
				t.Errorf("MixedCategory.CalcScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("MixedCategory.CalcScore() = %v, want %v", got, tt.expected)
			}
			if err == nil && mc.Used == false {
				t.Errorf("category should be marked used but is not")
			}
		})
	}
}

func TestFreeCategory_CalcScore(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		berries  []Berry
		expected int
		used bool
		wantErr  bool
	}{
		{
			name:     "Valid case - All different berries",
			berries:  []Berry{Jumbleberry, Sugarberry, Pickleberry, Moonberry, Pest},
			expected: 15, // 2 + 2 + 4 + 7 = 15 points (Pest counts as 0)
			wantErr:  false,
		},
		{
			name:     "Valid case - Only Jumbleberries",
			berries:  []Berry{Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry, Jumbleberry},
			expected: 10, // 5 * 2 points = 10 points
			wantErr:  false,
		},
		{
			name:     "Valid case - Only Sugarberries",
			berries:  []Berry{Sugarberry, Sugarberry, Sugarberry, Sugarberry, Sugarberry},
			expected: 10, // 5 * 2 points = 10 points
			wantErr:  false,
		},
		{
			name:     "Valid case - Mixed berries with Pest",
			berries:  []Berry{Jumbleberry, Jumbleberry, Sugarberry, Pickleberry, Pest},
			expected: 10, // 2 + 2 + 2 + 4 = 13 points (Pest counts as 0)
			wantErr:  false,
		},
		{
			name:     "Valid case - All berries are Pest",
			berries:  []Berry{Pest, Pest, Pest, Pest, Pest},
			expected: 0, // Pests count as 0 points
			wantErr:  false,
		},
		{
			name:     "Valid case - Pests and Jumbleberries",
			berries:  []Berry{Jumbleberry, Jumbleberry, Pest, Pest, Pest},
			expected: 4, // 2 + 2 = 4 points (Pests count as 0)
			wantErr:  false,
		},
		{
			name:     "Invalid case - Not exactly 5 berries",
			berries:  []Berry{Jumbleberry, Sugarberry, Pickleberry},
			expected: -1, // Incorrect number of berries
			wantErr:  true,
		},
		{
			name:     "Invalid case - More than 5 berries",
			berries:  []Berry{Jumbleberry, Sugarberry, Pickleberry, Moonberry, Pest, Pickleberry},
			expected: -1, // Incorrect number of berries
			wantErr:  true,
		},
		{
			name:     "Invalid case - Already Used",
			berries:  []Berry{Sugarberry, Pickleberry, Moonberry, Moonberry, Pest},
			expected: -1,
			used:     true,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fc := &FreeCategory{
				BaseCategory: BaseCategory{
					Used: tt.used,
				},
			}
			got, err := fc.CalcScore(tt.berries)
			if (err != nil) != tt.wantErr {
				t.Errorf("FreeCategory.CalcScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("FreeCategory.CalcScore() = %v, want %v", got, tt.expected)
			}
			if err == nil && fc.Used == false {
				t.Errorf("category should be marked used but is not")
			}
		})
	}
}
