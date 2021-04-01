package automorphic

import "testing"

func TestIsAutomorphicNumber(t *testing.T) {
	type testCase struct {
		name   string
		number int
		want   bool
	}
	tests := []testCase{
		{
			name:   "testing automorphic number 25",
			number: 25,
			want:   true,
		},
		{
			name:   "testing automorphic number 625",
			number: 625,
			want:   true,
		},
		{
			name:   "testing not automorphic number 13",
			number: 13,
			want:   false,
		},
		{
			name:   "testing not automorphic number 456",
			number: 456,
			want:   false,
		},
		{
			name:   "testing automorphic number 0",
			number: 0,
			want:   true,
		},
		{
			name:   "testing negative automorphic number -76",
			number: -76,
			want:   false,
		},
		{
			name:   "testing negative non automorphic number -13",
			number: -76,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumberAutomorphic(tt.number); got != tt.want {
				t.Errorf("result of automorphicity checking: %v, is number automorphic realy: %v", got, tt.want)
			}
		})
	}
}
