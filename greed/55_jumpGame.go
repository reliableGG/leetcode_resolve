package main

func canJump(nums []int) bool {

	if len(nums) == 0 {
		return true
	}
	maxReach := nums[0]
	for i := 1; i < len(nums) && i <= maxReach; i++ {
		maxReach = max(nums[i]+i, maxReach)
	}
	return maxReach >= len(nums)-1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
