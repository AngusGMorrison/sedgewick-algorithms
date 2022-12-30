package q4_dedup

import "sort"

func Dedup(s []string) []string {
	out := make([]string, len(s))
	copy(out, s)
	sort.Strings(out)

	i := 0
	j := i + 1
	for i < len(s)-1 {
		for ; j < len(s); j++ {
			if out[j] != out[i] {
				break
			}
		}

		if j == len(s) {
			break
		}

		i++
		out[i] = out[j]
	}

	return out[:i+1]
}
