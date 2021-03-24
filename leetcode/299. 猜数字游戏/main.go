package main

import "fmt"

func getHint(secret string, guess string) string {
	var same,shash,ghash = map[uint8]int{},map[uint8]int{},map[uint8]int{}

	for i:=0;i<len(secret);i++{
		s,g := secret[i],guess[i]

		if s == g {
			same[s]++
			continue
		}

		shash[s]++
		ghash[g]++
	}

	var countA int
	for _,num := range same{
		countA += num
	}

	var countB int

	for s,num := range shash{
		if gnum,ok := ghash[s];ok{
			if num > gnum {
				num = gnum
			}

			countB += num
		}
	}
	return fmt.Sprintf("%dA%dB", countA, countB)
}
