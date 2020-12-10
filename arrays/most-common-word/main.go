package main

import "fmt"

func mostCommonWord(paragraph string, banned []string) string {
	bannedTable := map[string]bool{}
	for _, bw := range banned {
		bannedTable[bw] = true
	}

	var ans string
	maxCount := 0

	wc := make(map[string]int)
	w := ""
	for i, r := range paragraph {
		if r != ' ' && r != ',' && r != '.' && r != '?' && r != ';' && r != '!' && r != '\'' {
			if r < 92 {
				r = r + 32
			}
			w = w + string(r)
		}
		if r == ' ' || r == ',' || r == '.' || r == '?' || r == ';' || r == '!' || (i == len(paragraph)-1) {
			_, bok := bannedTable[w]
			if bok {
				w = ""
				continue
			}
			if w == "" {
				continue
			}
			ct, ok := wc[w]

			ct++
			wc[w] = ct
			if ok {
				if ct > maxCount {
					ans = w
					maxCount = ct
				}
			} else {
				if maxCount == 0 {
					maxCount = 1
					ans = w
				}
			}
			w = ""
		}
	}

	if ans != "" {
		return ans
	}
	return w
}

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	fmt.Println(mostCommonWord("Bob", []string{}))
	fmt.Println(mostCommonWord("Bob. hIt, baLl", []string{"bob", "hit"}))
	fmt.Println(mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"}))
}
