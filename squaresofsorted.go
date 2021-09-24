func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
func sortedSquares(arr []int) []int {
	pop := 0
	hasN := false
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			hasN = true
		}
		if arr[i] >= 0 {
			pop = i
			break
		}
	}
	fmt.Print(pop)

	for x := 0; x < len(arr); x++ {
		arr[x] = arr[x] * arr[x]
	}
	fmt.Print(arr)
	if pop == 0 && hasN {
		return reverse(arr)
	}

	i, j := pop, pop-1
	var res []int
	// for m:=0;m<pop;m++ {
	//     arr[m]=arr[m]*-1
	// }
	for i < len(arr) && j >= 0 {
		if arr[i] <= arr[j] {
			res = append(res, arr[i])
			i++
		} else if arr[i] > arr[j] {
			res = append(res, arr[j])
			j--
		}
	}

	for i < len(arr) {
		res = append(res, arr[i])
		i++
	}
	for j >= 0 {
		res = append(res, arr[j])
		j--
	}

	return res
}