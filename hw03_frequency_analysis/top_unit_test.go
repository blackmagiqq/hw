package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWordsMatches(t *testing.T) {
	slice := []string{
		"a", "b", "c", "a",
	}

	expected := []wordCount{
		{
			word:  "a",
			count: 2,
		},
		{
			word:  "b",
			count: 1,
		},
		{
			word:  "c",
			count: 1,
		},
	}

	result := countWordsMatches(slice)

	for i, v := range expected {
		assert.Equal(t, v, result[i], "not equal")
	}
}

func TestSortingWordsCounts(t *testing.T) {
	inputStruct := []wordCount{
		{
			word:  "a",
			count: 2,
		},
		{
			word:  "b",
			count: 1,
		},
		{
			word:  "c",
			count: 1,
		},
	}

	expected := []string{
		"a", "b", "c",
	}

	result := sortingWordsCounts(inputStruct)

	for i, v := range expected {
		assert.Equal(t, v, result[i], "not equal")
	}
}
