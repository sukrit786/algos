func rotate(nums []int, k int) {

	n := len(nums)
	rot_arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print((i + k) % n)
		rot_arr[(i+k)%n] = nums[i]
	}
	fmt.Print(rot_arr)
	copy(nums, rot_arr)
}