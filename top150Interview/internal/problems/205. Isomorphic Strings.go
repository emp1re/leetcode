package problems

func IsIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	return checkPattern(s, t) && checkPattern(t, s)

}
func checkPattern(a string, b string) bool {
	hmap := make(map[string]string)
	for i := 0; i < len(a); i++ {
		value, ok := hmap[string(a[i])]
		if !ok {
			hmap[string(a[i])] = string(b[i])
		} else {
			if value != string(b[i]) {
				return false
			}
			hmap[string(a[i])] = string(b[i])
		}
	}
	return true
}
