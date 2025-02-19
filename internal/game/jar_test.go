package game

import (
	"testing"
)

// note that this test will roll a jar up 100 times and make sure that the berry changes at least once
// there is a very small chance this test could fail
func TestJar_Roll_Unlocked(t *testing.T) {
	t.Parallel()
	type fields struct {
		Berry  Berry
		Locked bool
		Rolled bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Jar Has Not Been Rolled",
			fields: fields{
				Berry:  Moonberry,
				Locked: false,
				Rolled: false,
			},
		},
		{
			name: "Jar Has Not Been Rolled",
			fields: fields{
				Berry:  Moonberry,
				Locked: false,
				Rolled: true,
			},
		},
		{

			name: "Jar Is Locked but unrolled",
			fields: fields{
				Berry:  Berry(-1),
				Locked: true,
				Rolled: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			j := &Jar{
				Berry:  tt.fields.Berry,
				Locked: tt.fields.Locked,
				Rolled: tt.fields.Rolled,
			}

			berryChanged := false
			berry := j.Berry
			// rolling 100 times to make sure the berry changes at least once
			for range 100 {
				j.Roll()
				if j.Berry != berry {
					berryChanged = true
					break
				}
			}

			if !berryChanged {
				t.Errorf("berry didn't change in 100 rolls")
			}
		})
	}
}

func TestJar_Roll_Locked(t *testing.T) {
	t.Parallel()
	type fields struct {
		Berry  Berry
		Locked bool
		Rolled bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Jar Is Locked",
			fields: fields{
				Berry:  Moonberry,
				Locked: true,
				Rolled: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			j := &Jar{
				Berry:  tt.fields.Berry,
				Locked: tt.fields.Locked,
				Rolled: tt.fields.Rolled,
			}

			// rolling 100 times to make sure the berry doesn't change
			for range 100 {
				j.Roll()
				if j.Berry != tt.fields.Berry {
					t.Errorf("berry changed but jar was supposed to be locked")
				}
			}
		})
	}
}

func TestJar_Lock(t *testing.T) {
	t.Parallel()
	type fields struct {
		Berry  Berry
		Locked bool
		Rolled bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Jar is Unlocked",
			fields: fields{
				Locked: false,
			},
		},
		{
			name: "Jar is Locked",
			fields: fields{
				Locked: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			j := &Jar{
				Berry:  tt.fields.Berry,
				Locked: tt.fields.Locked,
				Rolled: tt.fields.Rolled,
			}
			j.Lock()

			if j.Locked != true {
				t.Errorf("jar was supposed to be locked but wasn't")
			}
		})
	}
}

func TestJar_Unlock(t *testing.T) {
	t.Parallel()
	type fields struct {
		Berry  Berry
		Locked bool
		Rolled bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Jar is Unlocked",
			fields: fields{
				Locked: false,
			},
		},
		{
			name: "Jar is Locked",
			fields: fields{
				Locked: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			j := &Jar{
				Berry:  tt.fields.Berry,
				Locked: tt.fields.Locked,
				Rolled: tt.fields.Rolled,
			}
			j.Unlock()

			if j.Locked != false {
				t.Errorf("jar should be unlocked but isn't")
			}
		})
	}
}

func TestJar_Reset(t *testing.T) {
	t.Parallel()
	type fields struct {
		Berry  Berry
		Locked bool
		Rolled bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Unlocked and Unrolled",
			fields: fields{
				Locked: false,
				Rolled: false,
			},
		},
		{
			name: "Unlocked and Rolled",
			fields: fields{
				Locked: false,
				Rolled: true,
			},
		},
		{
			name: "Locked and Unrolled",
			fields: fields{
				Locked: true,
				Rolled: false,
			},
		},
		{
			name: "Locked and Rolled",
			fields: fields{
				Locked: true,
				Rolled: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			j := &Jar{
				Berry:  tt.fields.Berry,
				Locked: tt.fields.Locked,
				Rolled: tt.fields.Rolled,
			}
			j.Reset()

			if j.Rolled || j.Locked {
				t.Errorf("jar should be unrolled and unlocked but got unlocked: %t and rolled: %t", j.Locked, j.Rolled)
			}
		})
	}
}

func TestJar_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		Berry  Berry
		Locked bool
		Rolled bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Jumbleberry Unlocked and Rolled",
			fields: fields{
				Berry: Jumbleberry,
				Locked: false,
				Rolled: true,
			},
			want: Jumbleberry.String() + "✅",
		},
		{
			name: "Moonberry Locked and Rolled",
			fields: fields{
				Berry: Moonberry,
				Locked: true,
				Rolled: true,
			},
			want: Moonberry.String() + "🛑",
		},
		{
			name: "Unrolled Jar",
			fields: fields{
				Rolled: false,
			},
			want: "EMPTY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			j := &Jar{
				Berry:  tt.fields.Berry,
				Locked: tt.fields.Locked,
				Rolled: tt.fields.Rolled,
			}
			if got := j.String(); got != tt.want {
				t.Errorf("Jar.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
