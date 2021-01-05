package main

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		_, found := m[r]
		if !found {
			m[r] = 1
		} else {
			m[r] = m[r] + 1
		}
	}

	if len(m) == 0 {
		return -1
	}
	for i, r := range s {
		count, _ := m[r]
		if count == 1 {
			return i
		}
	}
	return -1
}

func main() {

}
