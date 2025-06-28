// Two pointers: one input, opposite ends

func twoPointerTemplate(arr []int) int {
    left, right := 0, len(arr)-1
    ans := 0

    for left < right {
        // TODO: logic involving arr[left] and arr[right]

        if /* some condition */ {
            left++
        } else {
            right--
        }
    }

    return ans
}


// Example usage

//https://leetcode.com/problems/container-with-most-water/description/

//https://leetcode.com/problems/valid-palindrome/submissions/1678915944/

//example

// func isPalindrome(s string) bool {
// 	left, right := 0, len(s)-1

// 	for left < right {
// 		for left < right && isWord(rune(s[left])) == false {
// 			left++
// 		}

// 		for left < right && isWord(rune(s[right])) == false {
// 			right--
// 		}

//         if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
//             return false
//         }

//         left ++
//         right--
// 	}

//     return true
// }

// func isWord(r rune) bool {
// 	return unicode.IsLetter(r) || unicode.IsDigit(r)
// }


// https://leetcode.com/problems/container-with-most-water/

// func maxArea(height []int) int {
//     left , right := 0 , len(height)-1
//     maxArea := 0

//     for left < right{
//         h := min(height[left],height[right])
//         width := right -left
//         area := h * width

//         if area > maxArea{
//             maxArea = area
//         }

//         if height[left] < height[right]{
//             left++
//         } else{
//             right --
//         }
//     }

//     return maxArea


// }


// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }


//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

// func twoSum(numbers []int, target int) []int {
// 	if len(numbers) == 2 {
// 		return []int{1, 2}
// 	}

// 	l, r := 0, len(numbers)-1

// 	for l < r {
//         sum := numbers[l] + numbers[r]
// 		if sum == target {
// 			return []int{l + 1, r + 1}
// 		}else if sum < target{
//             l++
//         } else{
//             r--
//         }
// 	}

//     return []int{}
// }

