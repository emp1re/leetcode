package problems

func UniqueMorseRepresentations(words []string) int {
	if len(words) == 1 {
		return 1
	}

	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	ab := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	dict := createDict(morse, ab)
	list := make([]string, len(words))

	for k := range words {
		for i := 0; i < len(words[k]); i++ {
			list[k] += string(dict[string(words[k][i])])
		}
	}

	hmap := map[string]int{}
	for i := 0; i < len(list); i++ {
		hmap[list[i]] += 1
	}

	return len(hmap)
}
func createDict(word []string, alphabet []string) (dict map[string]string) {

	dict = map[string]string{}
	for k, v := range word {
		dict[alphabet[k]] = v
	}
	return dict

}
