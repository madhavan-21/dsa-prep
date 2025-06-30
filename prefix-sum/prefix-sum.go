//Prefix Sum Template


func buildPrefixSum(arr []int) []int {
    n := len(arr)
    prefix := make([]int, n+1) // Extra space to simplify range queries
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + arr[i]
    }
    return prefix
}

// Get sum of arr[l..r] inclusive using prefix sum
func rangeSum(prefix []int, l, r int) int {
    return prefix[r+1] - prefix[l]
}

Example usage:
//https://leetcode.com/problems/range-sum-query-immutable/description/

// type NumArray struct {
//     nums []int
//     sums []int
// }


// func Constructor(nums []int) NumArray {
//     sums := makeSums(nums)
//     return NumArray {
//         nums: nums,
//         sums: sums,
//     }
// }

// func makeSums(nums []int) []int {
//     sums := make([]int, len(nums)+1)
//     for i := 1; i < len(sums); i++ {
//         sums[i] = sums[i-1] + nums[i-1]
//     }
//     return sums 
// }


// func (this *NumArray) SumRange(left int, right int) int {
//     return this.sums[right+1] - this.sums[left]
// }


//https://leetcode.com/problems/range-sum-query-immutable/
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


//https://leetcode.com/problems/find-pivot-index/

// func pivotIndex(nums []int) int {
//   var sum ,leftSum int
//   for _, value := range nums{
//     sum+=value
//   }  
//   for i := range nums{
//     if leftSum==sum-leftSum-nums[i]{
//         return i
//     }
//     leftSum +=nums[i]
//   }

//   return -1
// }