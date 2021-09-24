func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var arr []int
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i + 1
	}
	for i := 0; i < len(nums); i++ {
		req := target - nums[i]
		if m[req] != 0 && m[req]-1 != i {
			arr = append(arr, i, m[req]-1)
			m[req] = 0
			return arr
		}
	}
	// fmt.Println(m);
	// arr = append(arr,0,1);
	return arr
}