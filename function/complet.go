package function

import (
	"strings"

	fynex "fyne.io/x/fyne/widget"
)

func Complete(input fynex.CompletionEntry, entry []string, tab []string) []string {
	GetArtists()
	for x := 0; x < len(Artists); x++ { // Iterate through the array and if the name matches, add it to the array to display

		word2 := make([]string, len(Artists[x].Name))
		for i, r := range Artists[x].Name {
			word2[i] = string(r)
		}
		word := make([]string, len(input.Text))
		for i := 0; i < len(input.Text); i++ {
			if len(word2) >= len(input.Text) { // To avoid errors and program stop, words with a length less than the requested word's length will not pass
				word[i] = word2[i]
			}
		}
		// Convert words to uppercase for uniformity
		WordToMaj(entry)
		WordToMaj(word)
		textEntry := strings.Join(entry, "")
		mot := strings.Join(word, "")
		// Compare words to check for a match
		if textEntry == mot {
			tab = append(tab, Artists[x].Name+"-Artists") // Add the name to the tab array which will contain only names matching the name
		}
		// Members
		for y := 0; y < len(Artists[x].Members); y++ { // Iterate through each group's member array, and if the name matches, add it

			word2 := make([]string, len(Artists[x].Members[y]))
			for i, r := range Artists[x].Members[y] {
				word2[i] = string(r)
			}
			word := make([]string, len(input.Text))
			for i := 0; i < len(input.Text); i++ {
				if len(word2) >= len(input.Text) { // To avoid errors and program stop, words with a length less than the requested word's length
					word[i] = word2[i]
				}
			}
			WordToMaj(entry)
			WordToMaj(word)
			textEntry := strings.Join(entry, "")
			mot := strings.Join(word, "")
			// fmt.Println(mot)

			if textEntry == mot {
				tab = append(tab, Artists[x].Members[y]+"-Member") // Add the name to the tab array which will contain only names matching the name
			}

		}

	}

	return tab // Return the array
}
