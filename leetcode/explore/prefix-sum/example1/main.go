package main

// https://leetcode.com/explore/interview/card/leetcodes-interview-crash-course-data-structures-and-algorithms/703/arraystrings/4503/
func answerQueries(nums []int, queries [][]int, limit int) []bool {
	preSum := make([]int, len(nums))
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		preSum[i] = tmp
	}

	ans := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		left := queries[i][0]
		right := queries[i][1]
		curr := preSum[right] - preSum[left] + nums[i]
		ans[i] = curr < limit
	}

	return ans
}
