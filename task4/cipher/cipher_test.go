package cipher

import "testing"

func TestCaesarDecoder(t *testing.T) {
	type testCase struct {
		name   string
		cipher string
		keys   []string
		want   string
	}
	tests := []testCase{
		{
			"Decode right encoded string with only characters in it and right order of keys",
			"Xadqy Ubegy ue euybxk pgyyk fqjf ar ftq bduzfuzs mzp fkbqeqffuzs uzpgefdk",
			[]string{"Lorem", "dummy", "text", "industry"},
			"Lorem Ipsum is simply dummy text of the printing and typesetting industry",
		},
		{
			"Decode right encoded string with only characters in it and with no right order of keys",
			"Xadqy Ubegy ue euybxk pgyyk fqjf ar ftq bduzfuzs mzp fkbqeqffuzs uzpgefdk",
			[]string{"industry", "dummy", "Lorem", "text"},
			"Lorem Ipsum is simply dummy text of the printing and typesetting industry",
		},
		{
			"Decode right encoded string with non characters symbols in it and with no right order of keys",
			"Xadqy Ubegy! ue euybxk pgyyk fqjf: ar ftq bduzfuzs !mzp! fkbqeqffuzs ...uzpgefdk",
			[]string{"industry", "dummy", "Lorem", "text"},
			"Lorem Ipsum! is simply dummy text: of the printing !and! typesetting ...industry",
		},
		{
			"Decode right encoded string with wrong keys",
			"Xadqy Ubegy! ue euybxk pgyyk fqjf: ar ftq bduzfuzs !mzp! fkbqeqffuzs ...uzpgefdk",
			[]string{"test", "red", "black", "green"},
			"Failed to decrypt",
		},
		{
			"Decode non encoded string",
			"Lorem Ipsum is simply dummy text of the printing and typesetting industry",
			[]string{"Lorem", "dummy", "text", "industry"},
			"Lorem Ipsum is simply dummy text of the printing and typesetting industry",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaesarDecoder(tt.cipher, tt.keys...); got != tt.want {
				t.Errorf("String after decoding: %s, but expected: %s", got, tt.want)
			}
		})
	}
}
