package main

import "fmt"

var lt = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	6:    "VI",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	20:   "XX",
	30:   "XXX",
	40:   "XL",
	50:   "L",
	60:   "LX",
	70:   "LXX",
	80:   "LXXX",
	90:   "XC",
	100:  "C",
	200:  "CC",
	300:  "CCC",
	400:  "CD",
	500:  "D",
	600:  "DC",
	700:  "DCC",
	800:  "DCCC",
	900:  "CM",
	1000: "M",
	2000: "MM",
	3000: "MMM",
}

func intToRoman(num int) string {
	value, ok := lt[num]
	if ok {
		return value
	}
	return recurse(num, 1)
}

func recurse(number int, multiplier int) string {
	r := number % 10
	if number/10 <= 0 {
		return lt[r*multiplier]
	}

	return recurse(number/10, multiplier*10) + lt[r*multiplier]
}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(10))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(101))
	fmt.Println(intToRoman(1994))

}
