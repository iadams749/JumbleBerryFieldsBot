package game

import "testing"

func TestBerry_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		berry    Berry
		expected string
	}{
		{
			name:     "Pest Berry",
			berry:    Pest,
			expected: grey + "PEST" + reset,
		},
		{
			name:     "Jumbleberry",
			berry:    Jumbleberry,
			expected: red + "JBRY" + reset,
		},
		{
			name:     "Sugarberry",
			berry:    Sugarberry,
			expected: yellow + "SBRY" + reset,
		},
		{
			name:     "Pickleberry",
			berry:    Pickleberry,
			expected: green + "PBRY" + reset,
		},
		{
			name:     "Moonberry",
			berry:    Moonberry,
			expected: purple + "MBRY" + reset,
		},
		{
			name:     "Unknown Berry",
			berry:    Berry(999), // Invalid berry type
			expected: "Unknown Berry",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.berry.String()
			if got != tt.expected {
				t.Errorf("Berry.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

