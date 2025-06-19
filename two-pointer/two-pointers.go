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

