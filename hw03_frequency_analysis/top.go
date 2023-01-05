package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func Top10(textForAnalysis string) []string {
	words := strings.Fields(textForAnalysis)

	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}

	wordCountsSlice := []wordCount{}
	for word, count := range wordCounts {
		wordCountsSlice = append(wordCountsSlice, wordCount{word, count})
	}

	sort.Slice(wordCountsSlice, func(i, j int) bool {
		if wordCountsSlice[i].count == wordCountsSlice[j].count {
			return wordCountsSlice[i].word < wordCountsSlice[j].word
		}
		return wordCountsSlice[i].count > wordCountsSlice[j].count
	})

	var topWords []string
	for i := 0; i < 10 && i < len(wordCountsSlice); i++ {
		topWords = append(topWords, wordCountsSlice[i].word)
	}
	return topWords
}
