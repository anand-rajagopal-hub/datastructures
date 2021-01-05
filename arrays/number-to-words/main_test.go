package main

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		num  int
	}{
		{
			desc: "Zero",
			num:  0,
		},
		{
			desc: "One",
			num:  1,
		},
		{
			desc: "Twelve",
			num:  12,
		},
		{
			desc: "Ninety Nine",
			num:  99,
		},
		{
			desc: "One Hundred",
			num:  100,
		},
		{
			desc: "Three Hundred",
			num:  300,
		},
		{
			desc: "Three Hundred Forty Five",
			num:  345,
		},
		{
			desc: "Nine Hundred Ninety Nine",
			num:  999,
		},
		{
			desc: "One Thousand",
			num:  1000,
		},
		{
			desc: "Ten Thousand",
			num:  10000,
		},
		{
			desc: "Twelve Thousand",
			num:  12000,
		},
		{
			desc: "One Thousand Three Hundred Twenty Two",
			num:  1322,
		},
		{
			desc: "One Hundred Thousand",
			num:  100000,
		},
		{
			desc: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
			num:  1234567,
		},
		{
			desc: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
			num:  1234567891,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := numberToWords(tC.num)
			if strings.TrimSpace(actual) != tC.desc {
				t.Errorf("Test case failed for input %d. Expected = %s, Actual = %s ", tC.num, tC.desc, actual)
			}
		})
	}
}
