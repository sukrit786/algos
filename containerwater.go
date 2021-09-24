func min(i int, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
func maxArea(arr []int) int {
	i := 0
	j := len(arr) - 1
	max_area := 0
	for i < j {
		area := (j - i) * min(arr[i], arr[j])
		if arr[i] > arr[j] {
			j--
		} else {
			i++
		}
		if area > max_area {
			max_area = area
		}
	}
	return max_area
}