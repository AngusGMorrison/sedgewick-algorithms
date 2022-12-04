package q6_string_rotation

import "strings"

func isRotation(s1, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}
