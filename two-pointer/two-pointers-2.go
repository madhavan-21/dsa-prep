//Two pointers: two inputs, exhaust both


func fn(arr1, arr2 []int) int {
	i, j := 0, 0
	ans := 0

	for i < len(arr1) && j < len(arr2) {
		// do some logic
		if CONDITION {
			i++
		} else {
			j++
		}
	}

	for i < len(arr1) {
		// handle leftover arr1
		i++
	}

	for j < len(arr2) {
		// handle leftover arr2
		j++
	}

	return ans
}


// Example problems
//1 . https://leetcode.com/problems/merge-sorted-array/description/
//2 . https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
//3 . https://leetcode.com/problems/squares-of-a-sorted-array/description/
