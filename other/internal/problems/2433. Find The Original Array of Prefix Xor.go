package problems

func FindArray(pref []int) []int {
	out := make([]int, len(pref))
	for i := 0; i < len(pref); i++ {
		if i == 0 {
			out[i] ^= pref[i]
		} else {
			out[i] = pref[i] ^ pref[i-1]
		}
	}
	return out
}
