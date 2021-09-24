func moveZeroes(nums []int) {
	//     arr := make([]int, len(nums))

	//     n := len(nums);
	//     k:=0
	//     for i:=0;i<n;i++ {

	//         if(nums[i]!=0) {
	//             arr[k] = nums[i]
	//             k++
	//         }
	//     }

	//     copy(nums,arr)
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
		fmt.Print(nums)
	}

}