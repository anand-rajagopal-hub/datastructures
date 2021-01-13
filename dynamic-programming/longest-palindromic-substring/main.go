package main

import (
	"fmt"
)

func expandAroundCenter(b []byte) string {
	longest := 1
	longestStart := 0
	for i := 0; i < len(b)-1; i++ {
		if i > 0 && b[i-1] == b[i+1] {
			l := expand(b, i-1, i+1, 1)
			if l > longest {
				longest = l
				longestStart = i - l/2
			}
		}
		if b[i] == b[i+1] {
			l := expand(b, i-1, i+2, 2)
			if l > longest {
				longest = l
				longestStart = i - l/2 + 1
			}
		}
	}
	return string(b[longestStart : longestStart+longest])
}

func expand(b []byte, i, j, length int) int {
	if i < 0 || j >= len(b) {
		return length
	}

	if b[i] != b[j] {
		return length
	}

	return expand(b, i-1, j+1, length+2)
}

func longestPalindrome(s string) string {
	b := []byte(s)
	if len(b) == 1 {
		return s
	}
	if len(b) == 2 {
		if b[0] == b[1] {
			return string(b)
		}
		return string(b[0])
	}
	return expandAroundCenter(b)
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("ac"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("aacabdkacaa"))
	fmt.Println(longestPalindrome("ccc"))
	fmt.Println(longestPalindrome("abb"))
	fmt.Println(longestPalindrome("babaddtattarrattatddetartrateedredividerb"))
	fmt.Println(longestPalindrome("xaabacxcabaaxcabaax"))
	fmt.Println(longestPalindrome("abcda"))
	fmt.Println(longestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	fmt.Println(longestPalindrome("raedvmtyxveocfyhluuozodpxlroyjcsfslqmjthsbxhteeinpmnejxxcsyjgugclkehagpemfrnqtrkiropblcqdboztxtsaxqnsktrhzelynbzkxcghhfntrdauyzhzgujhniazijshesigzckgxentepeohcltpydumougjkmgoscchqsryaiamoujnyfpcsbwqtwikedbsjxxtnrpfgeqymwfngixslwlifimdapgzanuqwhwpesaigeoiwoyxzjmxukbsvsjvnjhwdbqzuurfolcngefdpsewrpvwivrsjnttrewkytdvvguatidyemrswpdmeqjrxgfdmcdlrcgiqdkyaaykdqigcrldcmdfgxrjqemdpwsrmeyditaugvvdtykwerttnjsrviwvprwespdfegnclofruuzqbdwhjnvjsvsbkuxmjzxyowioegiasepwhwqunazgpadmifilwlsxignfwmyqegfprntxxjsbdekiwtqwbscpfynjuomaiayrsqhccsogmkjguomudyptlchoepetnexgkczgisehsjizainhjugzhzyuadrtnfhhgcxkzbnylezhrtksnqxastxtzobdqclbporikrtqnrfmepgaheklcgugjyscxxjenmpnieethxbshtjmqlsfscjyorlxpdozouulhyfcoevxytmvdear"))
	fmt.Println(longestPalindrome("ccd"))

}
