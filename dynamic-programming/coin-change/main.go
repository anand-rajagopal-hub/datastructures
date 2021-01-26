package main

import (
	"fmt"
)

var memo = map[int]int{}

func cc(coins []int, amount int) int {
	v, ok := memo[amount]
	if ok {
		return v
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	result := amount + 1
	for _, v := range coins {
		ta := amount - v
		t := 0
		if ta > 0 {
			t = cc(coins, ta)
		}
		if t >= 0 {
			if result > t {
				result = t + 1
			}
		}

	}
	if result > amount {
		memo[amount] = -1
		return -1
	}
	memo[amount] = result
	return result
}

func coinChange(coins []int, amount int) int {
	memo = map[int]int{}

	return cc(coins, amount)
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange([]int{1}, 2))
	fmt.Println(coinChange([]int{1}, 1))
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	// fmt.Println(coinChange([]int{186, 83, 408}, 6249))
	// fmt.Println(coinChange([]int{419, 408}, 6249))
	// fmt.Println(coinChange([]int{83, 186, 408}, 383))

}
