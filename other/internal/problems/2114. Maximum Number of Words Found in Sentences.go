package problems

func MostWordsFound(sentences []string) int {
	var count int
	var max int
	for k := range sentences {
		for i := 0; i < len(sentences[k]); i++ {
			if string(sentences[k][i]) == " " {
				count++
			}
		}
		if count > max {
			max = count
		}
		count = 0
	}
	return max + 1
}
