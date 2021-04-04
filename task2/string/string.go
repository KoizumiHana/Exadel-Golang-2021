package string

import (
	"fmt"
	"strconv"
	"strings"
)

func Compress(stringToCompress string) string {
	count := 1
	prevChar := stringToCompress[0]
	var charsToReplace []string
	for i := 1; i < len(stringToCompress); i++ {
		if prevChar != stringToCompress[i] {
			if count > 3 {
				charsToReplace = append(charsToReplace, strings.Repeat(string(prevChar), count))
			}
			count = 1
			prevChar = stringToCompress[i]
		} else {
			count++
			if len(stringToCompress) == i+1 {
				charsToReplace = append(charsToReplace, strings.Repeat(string(prevChar), count))
			}
		}
	}
	for _, val := range charsToReplace {
		stringToCompress = strings.Replace(stringToCompress, val, fmt.Sprintf("#%v#%c", len(val), val[0]), 1)
	}
	return stringToCompress
}

func Decompress(stringToDecompress string) string {
	var decompressedString []byte
	for i := 0; i < len(stringToDecompress); i++ {
		if stringToDecompress[i] == '#' {
			if len(stringToDecompress[i:]) > 3 {
				if isReplaceableSequence(stringToDecompress[i : i+4]) {
					numberOfRepeating, _ := strconv.Atoi(string(stringToDecompress[i+1]))
					charToRepeat := string(stringToDecompress[i+3])
					decompressedCharSequence := []byte(strings.Repeat(charToRepeat, numberOfRepeating))
					decompressedString = append(decompressedString, decompressedCharSequence...)
					i += 3
				} else {
					decompressedString = append(decompressedString, stringToDecompress[i])
				}
			} else {
				decompressedString = append(decompressedString, []byte(stringToDecompress[i:])...)
				break
			}
		} else {
			decompressedString = append(decompressedString, stringToDecompress[i])
		}
	}
	return string(decompressedString)
}

func isReplaceableSequence(sequence string) bool {
	hasTwoHashSymbols := sequence[0] == '#' && sequence[2] == '#'
	_, err := strconv.Atoi(string(sequence[1]))
	return hasTwoHashSymbols && err == nil
}
