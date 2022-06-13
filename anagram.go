package main

func anagram(t string, s string) bool {
	m := map[rune]int{}

	if len(t) != len(s) {
		return false
	}

	for _, r := range s {
		m[r] = m[r] + 1
	}

	for _, r := range t {
		if m[r] == 0 {
			return false
		}
		m[r] = m[r] - 1
	}

	return true
}
