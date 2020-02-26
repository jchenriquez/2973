package mostcommonword

import ("fmt"
"unicode"
)

func MostCommonWord(paragraph string, banned []string) string {
  wordEndings := map[byte]bool {
    '.': true,
    ',': true,
    ';': true,
    '!': true,
    ' ': true,
    '\t': true,
    '\n': true,
  }

  bannedLkup := make(map[string] bool, len(banned))

  for _, ban := range banned {
    bannedLkup[ban] = true
  }

  var maxWord string
  var maxWordCount int
  wordsCount := make(map[string]int, 10)
  var currWord string

  for i := range paragraph {
    letter := paragraph[i]

    if _, wordEnded := wordEndings[letter]; wordEnded {
      if _, baned := bannedLkup[currWord]; len(currWord) > 0 && !baned {
        wordCount := wordsCount[currWord]
        wordsCount[currWord] = wordCount+1

        if wordCount+1 > maxWordCount {
          maxWordCount = wordCount+1
          maxWord = currWord
        }

      } 

      currWord = ""
    } else {
     currWord = fmt.Sprintf("%s%c", currWord, unicode.Lower(letter))
    }
  }

  return maxWord
}