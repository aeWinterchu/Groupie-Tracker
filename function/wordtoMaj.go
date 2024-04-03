package function

import (
	"strings"
)

func WordToMaj(word_tab []string) {
	// Convert the word array to a string, then to a rune array
	mot := strings.Join(word_tab, "")
	runes := []rune(mot)
	var result []int
	// Traverse the rune array and convert lowercase letters to uppercase while respecting characters with the most frequent accents
	for i := 0; i <= len(runes)-1; i++ {
		result = append(result, int(runes[i]))
		if result[i] == 250 {
			result[i] = result[i] - 165
			word_tab[i] = string(result[i])

		} else if result[i] == 233 {
			result[i] = result[i] - 164
			word_tab[i] = string(result[i])
		} else if result[i] > 96 && result[i] < 123 {
			result[i] = result[i] - 32
			word_tab[i] = string(result[i])
		}
	}

}
