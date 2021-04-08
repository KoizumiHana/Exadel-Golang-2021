package cipher

import (
	"strings"
)

func CaesarDecoder(cipher string, keys ...string) string {
	var decoded string
	var rotate func(rune) rune
	for shift := 0; shift <= 25; shift++ {
		rotate = func(r rune) rune {
			switch {
			case r >= 'A' && r <= 'Z':
				return 'A' + (r-'A'+rune(shift))%26
			case r >= 'a' && r <= 'z':
				return 'a' + (r-'a'+rune(shift))%26
			}
			return r
		}
		decoded = strings.Map(rotate, cipher)
		if !containsSubstrings(decoded, keys...) {
			continue
		}
		return decoded
	}
	return "Failed to decrypt"
}

func containsSubstrings(str string, subs ...string) bool {
	isCompleteMatch := true
	for _, sub := range subs {
		if !strings.Contains(str, sub) {
			isCompleteMatch = false
		}
	}
	return isCompleteMatch
}
