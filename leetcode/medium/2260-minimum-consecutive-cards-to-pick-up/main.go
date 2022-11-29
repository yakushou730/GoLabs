package main

// https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/
func minimumCardPickup(cards []int) int {
	cardsMap := make(map[int][]int, 0)
	ans := 100_00_00
	for i, card := range cards {
		cardsMap[card] = append(cardsMap[card], i)
		if len(cardsMap[card]) < 2 {
			continue
		} else {
			diff := cardsMap[card][len(cardsMap[card])-1] - cardsMap[card][len(cardsMap[card])-2] + 1
			if diff < ans {
				ans = diff
			}
		}
	}

	if ans == 100_00_00 {
		return -1
	} else {
		return ans
	}
}
