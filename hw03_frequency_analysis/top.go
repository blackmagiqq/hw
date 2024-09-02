package hw03frequencyanalysis

import (
	"sort"
	"strings"

	"github.com/go-corelibs/slices"
)

type wordCount struct {
	word  string
	count int
}

func Top10(text string) []string {
	arrayOfWords := strings.Fields(text)
	sliceOfCountWordsMatches := countWordsMatches(arrayOfWords)
	sortedWordsSlice := sortingWordsCounts(sliceOfCountWordsMatches)

	return sortedWordsSlice
}

func countWordsMatches(s []string) []wordCount {
	result := []wordCount{}

	for _, v := range s {
		indexes := slices.IndexesOf(s, v)
		result = append(result, wordCount{v, len(indexes)})
		for i := len(indexes) - 1; i > 0; i-- {
			s = slices.Remove(s, indexes[i])
		}
	}
	return result
}

func sortingWordsCounts(s []wordCount) []string {
	result := []string{}
	sort.Slice(s, func(i, j int) bool {
		if s[i].count == s[j].count {
			return s[i].word < s[j].word
		}
		return s[i].count > s[j].count
	})

	for _, v := range s {
		if len(result) > 9 {
			break
		}
		result = append(result, v.word)
	}
	return result
}
