// sliding windows

func fn(arr []int) int {
	left, ans, curr := 0, 0, 0

	for right := 0; right < len(arr); right++ {
		// 1. Add arr[right] to the current window state
		// Example: curr += arr[right]

		// 2. While the window condition is broken, shrink it from the left
		for WINDOW_CONDITION_BROKEN {
			// Remove arr[left] from curr
			// Example: curr -= arr[left]
			left++
		}

		// 3. Update the result
		// Example: ans = max(ans, curr)
	}

	return ans
}

// Example problems
// 1. https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
// 2. https://leetcode.com/problems/longest-repeating-character-replacement
// 3. https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/description/


//1.https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(str string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < len(str); right++ {
		if index, found := charIndex[str[right]]; found && index >= left {
			left = index + 1
		}
		charIndex[str[right]] = right
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}