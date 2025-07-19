// 1423. Maximum Points You Can Obtain from Cards
// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/description/
// https://go.dev/play/p/QyW0dNB3sA6


// func main() {
// 	fmt.Println("Hello, 世界")
// 	cardpoints := []int{5, 6, 1, 1, 1, 1, 6}
// 	k := 3
// 	res := maxScore(cardpoints, k)
// 	fmt.Println("result===============>", res)
// }


// func maxScore(cardPoints []int, k int) int {
//     leftSum :=0
//     rightSum :=0
//     maxSum :=0

//     for i :=0;i <= k-1 ;i++{
//         leftSum = leftSum + cardPoints[i]
//     }

//     maxSum = max(maxSum , leftSum)
//     rightIndex := len(cardPoints) -1

//     for j:= k-1;j>=0;j--{
//         leftSum = leftSum-cardPoints[j]
//         rightSum = rightSum +cardPoints[rightIndex]
//         rightIndex--


//         maxSum = max(maxSum , leftSum+rightSum)
//     }

//     return maxSum
// }

//-------------------------------------------------------------------------------------------------------------------------------------------------------------------

// L3. Longest Substring Without Repeating Characters
// https://go.dev/play/p/cK5gmF8YS8Z
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// func main() {
// 	fmt.Println("Hello, 世界")
// 	str := "abcabcbb"

// 	longesSubstring := lengthOfLongestSubstring(str)
// 	fmt.Println("longestSubStrig=============>", longesSubstring)

// }

// func lengthOfLongestSubstring(str string) int {
// 	charIndex := make(map[byte]int)
// 	maxLen := 0
// 	left := 0

// 	for right := 0; right < len(str); right++ {
// 		if index, found := charIndex[str[right]]; found && index >= left {
// 			left = index + 1
// 		}

// 		charIndex[str[right]] = right
// 		maxLen = max(maxLen, right-left+1)
// 	}

// 	return maxLen
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

//---------------------------------------------------------------------------------------------------------------------------------------------------------

// 1004. Max Consecutive Ones III
// https://go.dev/play/p/ohF2gI_-Bld
// https://leetcode.com/problems/max-consecutive-ones-iii/description/

// func longestOnes(nums []int, k int) int {
// 	// left := 0
// 	// zeros := 0
// 	// maxlen := 0

// 	// for right := 0; right < len(nums); right++ {
// 	//     if nums[right] == 0 {
// 	//         zeros++
// 	//     }

// 	//     for zeros > k {
// 	//         if nums[left] == 0 {
// 	//             zeros--
// 	//         }
// 	//         left++
// 	//     }

// 	//     maxlen = max(maxlen, right-left+1)
// 	// }

// 	// return maxlen

// 	left := 0
// 	right := 0
// 	maxLen := 0
// 	zeros := 0
// 	for right < len(nums) {
// 		if nums[right] == 0 {
// 			zeros++
// 		}

// 		if zeros > k {
// 			if nums[left] == 0 {
// 				zeros--
// 			}
// 				left++
// 		}
// 		if zeros <= k {
// 			maxLen = max(maxLen, right-left+1)
// 		}
//         right++
// 	}

// 	return maxLen
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }


//---------------------------------------------------------------------------------------------------------------------------------------------------------


// 904. Fruit Into Baskets
// https://go.dev/play/p/BSzIVDZIvR3

// func main() {
// 	fmt.Println("Hello, 世界")
// 	fruits := []int{1, 2, 1}
// 	res := totalFruit(fruits)
// 	fmt.Println("response=============>", res)
// }

// func totalFruit(fruits []int) (res int) {
// 	var (
// 		buckets = make(map[int]int, 3)
// 		l, r    int
// 	)

// 	// for ;r <len(fruits); r++{
// 	//     buckets[fruits[r]]++
// 	//     for len(buckets) >2 {
// 	//         buckets[fruits[l]]--
// 	//         if buckets[fruits[l]]==0{
// 	//             delete(buckets , fruits[l])
// 	//         }
// 	//         l++
// 	//     }
// 	//     res = max(res , r-l+1)
// 	// }

// 	for r < len(fruits) {
// 		buckets[fruits[r]]++

// 		if len(buckets) > 2 {
// 			buckets[fruits[l]]--

// 			if buckets[fruits[l]] == 0 {
// 				delete(buckets, fruits[l])
// 			}
// 			l++
// 		}
// 		res = max(res, r-l+1)
// 		r++
// 	}
// 	return res
// }

//----------------------------------------------------------------------------------------------------------------------------------------------------------

// 1358. Number of Substrings Containing All Three Characters
//  https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/description/
// https://go.dev/play/p/SrCY9XTjS91


// You can edit this code!
// Click here and start typing.
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界")
// 	s := "abcabc"
// 	res := numberOfSubstrings(s)
// 	fmt.Println(res)
// }

// func numberOfSubstrings(s string) int {
// 	last := [3]int{-1, -1, -1}
// 	ans := 0

// 	for i := 0; i < len(s); i++ {
// 		last[s[i]-'a'] = i
// 		if last[0] != -1 && last[1] != -1 && last[2] != -1 {
// 			ans += 1 + min(last[0], last[1], last[2])
// 		}
// 	}

// 	return ans
// }

// func min(a, b, c int) int {
// 	if a < b {
// 		if a < c {
// 			return a
// 		}
// 		return c
// 	}
// 	if b < c {
// 		return b
// 	}
// 	return c
// }


//----------------------------------------------------------------------------------------------------------------------------------------------------------

// 424. Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/description/
// https://go.dev/play/p/eVDdr7aysn1

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界")
// 	s := "ABAB"
// 	k := 2
// 	res := characterReplacement(s, k)
// 	fmt.Println("response=========>", res)
// }

// func characterReplacement(s string, k int) int {
// 	left := 0
// 	right := 0
// 	maxLen := 0
// 	maxRepeat := 0
// 	freq := make(map[byte]int)

// 	for right < len(s) {
// 		freq[s[right]]++
// 		maxRepeat = max(maxRepeat, freq[s[right]])

// 		if (right-left+1)-maxRepeat > k {
// 			freq[s[left]]--
// 			left++
// 		}

// 		if (right-left+1)-maxRepeat <= k {
// 			maxLen = max(maxLen, right-left+1)
// 		}

// 		right++
// 	}

// 	return maxLen
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

//----------------------------------------------------------------------------------------------------------------------------------------------------------

// 930. Binary Subarrays With Sum

// https://leetcode.com/problems/binary-subarrays-with-sum/description/
//https://go.dev/play/p/b2xF3dsB3x9

// You can edit this code!
// Click here and start typing.
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界")
// 	nums := []int{1, 0, 1, 0, 1}
// 	goal := 2
// 	res := numSubarraysWithSum(nums, goal)
// 	fmt.Println("res==========> ", res)
// }

// func numSubarraysWithSum(nums []int, goal int) int {
// 	first := sumAtMost(nums, goal)
// 	second := sumAtMost(nums, goal-1)
// 	fmt.Printf("first-------------------%v , second-------------------------%v \n", first, second)
// 	return sumAtMost(nums, goal) - sumAtMost(nums, goal-1)
// }

// func sumAtMost(nums []int, goal int) int {
// 	if goal < 0 {
// 		return 0
// 	}

// 	res, prev, cur := 0, 0, 0

// 	for i, val := range nums {
// 		cur += val
// 		for cur > goal {
// 			cur -= nums[prev]
// 			prev++
// 		}
// 		res += i - prev + 1
// 	}

// 	return res
// }


//----------------------------------------------------------------------------------------------------------------------------------------------------------\
// 1248. Count Number of Nice Subarrays
// https://leetcode.com/problems/count-number-of-nice-subarrays/description/
// https://go.dev/play/p/VgBiaCxtDOZ

// You can edit this code!
// Click here and start typing.
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界")
// 	nums := []int{1, 1, 2, 1, 1}
// 	k := 3

// 	res := numberOfSubarrays(nums, k)
// 	fmt.Println("res------------->", res)
// }

// func numberOfSubarrays(nums []int, k int) int {
// 	prefixCounts := map[int]int{0: 1}
// 	oddCount := 0
// 	count := 0

// 	for _, num := range nums {
// 		if num%2 != 0 {
// 			oddCount++
// 		}
// 		if v, ok := prefixCounts[oddCount-k]; ok {
// 			count += v
// 		}
// 		prefixCounts[oddCount]++
// 	}

// 	return count
// }



//-------------------------------------------------------------------------------------------------------------------------------------------------------
// 992. Subarrays with K Different Integers
// https://leetcode.com/problems/subarrays-with-k-different-integers/description/
// https://go.dev/play/p/Jws9gHmuJma

// func main() {
// 	fmt.Println("Hello, 世界")
// 	nums := []int{1, 2, 1, 2, 3}
// 	k := 2
// 	val := subarraysWithKDistinct(nums, k)
// 	fmt.Println("resp===========>", val)
// }

// func subarraysWithAtMostKDistinct(nums []int, k int) int {
// 	if k == 0 {
// 		return 0
// 	}

// 	countOccurrence := make(map[int]int)
// 	//differentIntegers := 0
// 	left := 0
// 	result := 0

// 	for right := 0; right < len(nums); right++ {
// 		countOccurrence[nums[right]]++
// 		// if countOccurrence[nums[right]]==1{
// 		//     differentIntegers++
// 		// }

// 		for len(countOccurrence) > k {
// 			countOccurrence[nums[left]]--
// 			if countOccurrence[nums[left]] == 0 {
// 				//differentIntegers--
// 				delete(countOccurrence, nums[left])
// 			}
// 			left++
// 		}

// 		result += right - left + 1
// 	}

// 	return result
// }

// func subarraysWithKDistinct(nums []int, k int) int {
// 	return subarraysWithAtMostKDistinct(nums, k) - subarraysWithAtMostKDistinct(nums, k-1)
// }

//------------------------------------------------------------------------------------------

// 76. Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/description/
// https://go.dev/play/p/8nezFmeeRze

// You can edit this code!
// Click here and start typing.
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界")
// 	s := "ADOBECODEBANC"
// 	t := "ABC"
// 	res := minWindow(s, t)
// 	fmt.Println("response===================>", res)
// }

// func minWindow(s string, t string) string {
// 	if len(t) == 0 || len(s) < len(t) {
// 		return ""
// 	}

// 	// Step 1: Count required characters in t
// 	need := make(map[byte]int)
// 	for i := 0; i < len(t); i++ {
// 		need[t[i]]++
// 	}

// 	// Step 2: Sliding window
// 	window := make(map[byte]int)
// 	have, needCount := 0, len(need)
// 	left := 0
// 	minLen := len(s) + 1
// 	start := 0

// 	for right := 0; right < len(s); right++ {
// 		char := s[right]
// 		window[char]++

// 		if _, ok := need[char]; ok && window[char] == need[char] {
// 			have++
// 		}

// 		// Step 3: Try to shrink the window
// 		for have == needCount {
// 			// Update result if smaller window found
// 			if right-left+1 < minLen {
// 				minLen = right - left + 1
// 				start = left
// 			}

// 			// Move left pointer
// 			leftChar := s[left]
// 			window[leftChar]--
// 			if _, ok := need[leftChar]; ok && window[leftChar] < need[leftChar] {
// 				have--
// 			}
// 			left++
// 		}
// 	}

// 	if minLen == len(s)+1 {
// 		return ""
// 	}

// 	return s[start : start+minLen]
// }

