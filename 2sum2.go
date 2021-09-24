func twoSum(arr []int, target int) []int {
	l, r := 0, len(arr)-1
	// out:=make([]int,5)
	out := []int{}
	for l < r {
		sum := arr[l] + arr[r]
		if sum == target {
			out = append(out, []int{l + 1, r + 1}...)
			break
		}
		if sum < target {
			l++
		}
		if sum > target {
			r--
		}

	}
	return out
}