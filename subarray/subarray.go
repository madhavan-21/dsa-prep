//Find number of subarrays that fit an exact criteria

func countSubarrays(arr []int, k int) int {
	counts := make(map[int]int)
	counts[0] = 1
	ans, curr := 0, 0

	for _, num := range arr {
		// Do logic to change curr
		// For example: curr += num   (for sum-based problems)
		// You should replace this line based on the specific problem
		curr += num  

		ans += counts[curr - k]
		counts[curr]++
	}

	return ans
}

//https://leetcode.com/problems/subarray-sum-equals-k/submissions/1685293361/
// func subarraySum(nums []int, k int) int {
// 	count := 0
// 	prefixSum := 0
// 	freq := map[int]int{0: 1}
// 	for i := 0; i < len(nums); i++ {
// 		prefixSum += nums[i]
// 		count += freq[prefixSum-k]
// 		freq[prefixSum] += 1
// 	}
// 	return count
// }


//https://leetcode.com/problems/continuous-subarray-sum/submissions/1685304293/

// func checkSubarraySum(nums []int, k int) bool {
//     remaindersFound := make(map[int]int)
//     remaindersFound[0] = -1

//     for i :=0 ; i < len(nums);i++{
//         currSum += nums[i]
//         remainder := currSum % k

//         if val, exists := remaindersFound[remainder]; exists {
//             // Check if the length of the subarray is at least 2
//             if i - val >= 2 {
//                 return true
//             }
//         } else {
//             // Store the remainder and the current index
//             remaindersFound[remainder] = i
//         }
//     }

//     return false
// }