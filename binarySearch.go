func search(nums []int, target int) int {
	for i, j := 0, len(nums)-1; i <= j; {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else if nums[m] > target {
			j = m - 1
		} else {
			return m
		}
	}
	return -1
}