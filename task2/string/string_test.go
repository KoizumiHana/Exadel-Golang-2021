package string

import (
	"testing"
)

type testCase struct {
	name   string
	string string
	want   string
}

func TestCompress(t *testing.T) {
	tests := []testCase{
		{
			"compression of string with enough repeated letters",
			"veeeeector",
			"v#5#ector",
		},
		{
			"compression of string with enough repeated letters but with space in between",
			"veee eector",
			"veee eector",
		},
		{
			"compression of the string without repeated letters",
			"vector",
			"vector",
		},
		{
			"compression of the short string",
			"dog",
			"dog",
		},
		{
			"compression of the string with less than 4 repeated letters",
			"veeector",
			"veeector",
		},
		{
			"compression of the string with multiple repeated letters",
			"veeeectoooooor",
			"v#4#ect#6#or",
		},
		{
			"compression of the string with repeated letters at the end",
			"vectorrrrr",
			"vecto#5#r",
		},
		{
			"compression of the string with repeated letters at the beginning",
			"vvvvvector",
			"#5#vector",
		},
		{
			"compression of the string with multiple repeated letters, but one sequence too short",
			"veeeectooor",
			"v#4#ectooor",
		},
		{
			"compression of the string with with only repeated letters in it",
			"oooooo",
			"#6#o",
		},
		{
			"compression of the string with with only repeated letters in it multiple times",
			"oooooorrrrr",
			"#6#o#5#r",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compress(tt.string); got != tt.want {
				t.Errorf("String after compression: %s, but expected: %s", got, tt.want)
			}
		})
	}
}

func TestDecompress(t *testing.T) {
	tests := []testCase{
		{
			"decompression of string with correct sequence of char for decompression",
			"v#5#ector",
			"veeeeector",
		},
		{
			"decompression of the string without sequence of char for decompression",
			"vector",
			"vector",
		},
		{
			"decompression of the string with multiple sequence of char for decompression",
			"v#4#ect#6#or",
			"veeeectoooooor",
		},
		{
			"decompression of the string with sequence of char for decompression at the end",
			"vecto#5#r",
			"vectorrrrr",
		},
		{
			"decompression of the string with sequence of char for decompression at the beginning",
			"#5#vector",
			"vvvvvector",
		},
		{
			"decompression of the string with broken sequence of char for decompression at the beginning",
			"#i#vector",
			"#i#vector",
		},
		{
			"decompression of the string with partial sequence of char for decompression at the beginning",
			"#5vector",
			"#5vector",
		},
		{
			"decompression of the string with broken sequence of char for decompression at the end",
			"vecto#r#r",
			"vecto#r#r",
		},
		{
			"decompression of the string with with only sequence of char for decompression in it",
			"#6#o",
			"oooooo",
		},
		{
			"decompression of the string with only part of char sequence for decompression in it",
			"#6#",
			"#6#",
		},
		{
			"decompression of the string with with only sequence of char for decompression in it multiple times",
			"#6#o#5#r",
			"oooooorrrrr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decompress(tt.string); got != tt.want {
				t.Errorf("String after decompression: %s, but expected: %s", got, tt.want)
			}
		})
	}
}
