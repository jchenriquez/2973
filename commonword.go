package mostcommonword

import (
	"fmt"
	"unicode"
)

// MostCommonWord will return the most non-banned word in parahraph
func MostCommonWord(paragraph string, banned []string) string {
	wordEndings := map[byte]bool{
		'.':  true,
		',':  true,
		';':  true,
		'!':  true,
		' ':  true,
		'\t': true,
		'\n': true,
		'?':  true,
		'\'': true,
	}

	bannedLkup := make(map[string]bool, len(banned))

	for _, ban := range banned {
		bannedLkup[ban] = true
	}

	var maxWord string
	var maxWordCount int
	wordsCount := make(map[string]int, 10)
	var currWord string

	for i := 0; i <= len(paragraph); i++ {
		var letter byte

		if i < len(paragraph) {
			letter = paragraph[i]
		}

		if _, wordEnded := wordEndings[letter]; wordEnded || i == len(paragraph) {
			if _, baned := bannedLkup[currWord]; len(currWord) > 0 && !baned {
				wordCount := wordsCount[currWord]
				wordsCount[currWord] = wordCount + 1

				if wordCount+1 > maxWordCount {
					maxWordCount = wordCount + 1
					maxWord = currWord
				}

			}

			currWord = ""
		} else {
			currWord = fmt.Sprintf("%s%c", currWord, unicode.ToLower(rune(letter)))
		}
	}

	return maxWord
}
