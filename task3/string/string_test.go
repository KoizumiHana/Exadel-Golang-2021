package string

import "testing"

func TestGroupByOccurrences(t *testing.T) {
	type testCase struct {
		name   string
		string string
		want   string
	}
	tests := []testCase{
		{
			"grouping of example string",
			"one, two - this 2, three one two, lot of words: one",
			"one(3) two(2) this(1) 2(1) three(1) lot(1) of(1) words(1)",
		},
		{
			"grouping of string without multiple occurrences of same words",
			"one, two - this 2, three, lot of words:",
			"one(1) two(1) this(1) 2(1) three(1) lot(1) of(1) words(1)",
		},
		{
			"grouping of example string with non consecutive indexes of words",
			"one, two - this 2, three one two, lot of words: one! words:",
			"one(3) two(2) words(2) this(1) 2(1) three(1) lot(1) of(1)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupByOccurrences(tt.string); got != tt.want {
				t.Errorf("String after groupping: %s, but expected: %s", got, tt.want)
			}
		})
	}
}
