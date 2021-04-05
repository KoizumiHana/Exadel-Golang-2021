package string

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func GroupByOccurrences(stringToGroup string) string {
	words := strings.FieldsFunc(
		stringToGroup,
		func(r rune) bool {
			return !unicode.IsLetter(r)
		})

	type oi struct {
		occurrences int
		index       int
	}

	wordsOccurrences := make(map[string]oi)

	for index, word := range words {
		if v, ok := wordsOccurrences[word]; ok {
			v.occurrences++
			wordsOccurrences[word] = v
		} else {
			wordsOccurrences[word] = oi{1, index}
		}
	}

	type woi struct {
		word        string
		occurrences int
		index       int
	}

	var wois []woi
	for k, v := range wordsOccurrences {
		wois = append(wois, woi{k, v.occurrences, v.index})
	}

	sort.Slice(wois, func(i, j int) bool {
		if wois[i].occurrences > wois[j].occurrences {
			return true
		}
		if wois[i].occurrences < wois[j].occurrences {
			return false
		}
		return wois[i].index < wois[j].index
	})

	var groupedString string
	for _, woi := range wois {
		groupedString += fmt.Sprintf("%s(%v) ", woi.word, woi.occurrences)
	}
	return strings.TrimSpace(groupedString)
}
