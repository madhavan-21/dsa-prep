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



// https://leetcode.com/problems/merge-sorted-array/

// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     p1 , p2 , p := m-1, n-1 , m+n-1

//     for p1 >= 0 && p2 >= 0{
//         if nums1[p1] > nums2[p2]{
//             nums1[p] = nums1[p1]
//             p1--

//         }else{
//             nums1[p] = nums2[p2]
//             p2--
//         }
//         p--
//     }

//     for p2 >= 0{
//         nums1[p] = nums2[p2]
//         p2--
//         p--
//     }
// }

//https://leetcode.com/problems/merge-two-sorted-lists/submissions/1678976443/

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	curr := dummy

// 	for l1 != nil && l2 != nil {
//         if l1.Val < l2.Val{
//             curr.Next = l1
//             l1 = l1.Next
//         }else{
//             curr.Next = l2
//             l2 = l2.Next
//         }
//         curr = curr.Next
// 	}

//     if l1 != nil{
//         curr.Next = l1
//     }else{
//         curr.Next = l2
//     }

//     return dummy.Next

// }


//https://leetcode.com/problems/intersection-of-two-arrays-ii/submissions/1678986079/

// func intersect(nums1 []int, nums2 []int) []int {
//    sort.Ints(nums1)
// 	sort.Ints(nums2)

//     n1 , n2 := 0 ,0

//     result := []int{}

//     for n1 < len(nums1) && n2 < len(nums2){
//         if nums1[n1] == nums2[n2]{
//             result = append(result , nums1[n1])
//             n1++
//             n2++
//         }else if nums1[n1] < nums2[n2]{
//            n1++
//         }else{
//             n2++
//         }
//     }

//     return result
// }
