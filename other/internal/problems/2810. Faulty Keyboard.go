package problems

func FinalString(s string) string {
	var str string
	for _, v := range s {
		if string(v) == "i" {
			var rev string
			for _, val := range str {
				rev = string(val) + rev
			}
			str = rev
		} else {
			str = str + string(v)
		}

	}
	return str
}
