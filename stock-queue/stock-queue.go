// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/description/
// https://go.dev/play/p/yCtekNEm9f0

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界")
// 	str := "()[]{}"
// 	isvalid := isValid(str)
// 	fmt.Println("isvalid parenthisis", isvalid)
// }

// func isValid(s string) bool {
// 	stack := []rune{}
// 	pairs := map[rune]rune{
// 		'}': '{',
// 		')': '(',
// 		']': '[',
// 	}

// 	for _, char := range s {
// 		if char == '(' || char == '{' || char == '[' {
// 			stack = append(stack, char)
// 		} else {
// 			if len(stack) == 0 {
// 				return false
// 			}
// 			if pairs[char] != stack[len(stack)-1] {
// 				return false
// 			}

// 			stack = stack[:len(stack)-1]
// 		}
// 	}

// 	return len(stack) == 0
// }


// 155. Min Stack
// https://leetcode.com/problems/min-stack/description/

// type MinStack struct {
//     Val []int
// }


// func Constructor() MinStack {
//     return MinStack{}
// }


// func (this *MinStack) Push(val int)  {
//     this.Val = append(this.Val , val)
// }


// func (this *MinStack) Pop()  {
//     this.Val = this.Val[:len(this.Val)-1]
// }


// func (this *MinStack) Top() int {
//     return this.Val[len(this.Val)-1]
// }


// func (this *MinStack) GetMin() int {
//     min := this.Val[0]
//     for _, v := range this.Val{
//         if v < min{
//             min = v
//         }
//     }

//     return min
// }


// 496. Next Greater Element I
// https://leetcode.com/problems/next-greater-element-i/description/
// https://go.dev/play/p/RPwZxYkjwmY

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界")
// 	nums1 := []int{4, 1, 2}
// 	nums2 := []int{1, 3, 4, 2}
// 	res := nextGreaterElement(nums1, nums2)
// 	fmt.Println("nexr greater elements============>", res)
// }

// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	nextGreater := map[int]int{}
// 	stack := []int{}

// 	for _, num := range nums2 {
// 		nextGreater[num] = -1

// 		for len(stack) > 0 && stack[len(stack)-1] < num {
// 			top := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			nextGreater[top] = num
// 		}

// 		stack = append(stack, num)
// 	}
// 	res := []int{}
// 	for _, num := range nums1 {
// 		res = append(res, nextGreater[num])
// 	}
// 	return res
// }


// 503. Next Greater Element II
// https://leetcode.com/problems/next-greater-element-ii/description/
// https://go.dev/play/p/fWfZaPALsYx

// func main() {
// 	fmt.Println("Hello, 世界")
// 	elements := []int{1, 2, 1}
// 	res := nextGreaterElements(elements)
// 	fmt.Println("res======================>", res)
// }

// func nextGreaterElements(nums []int) []int {
// 	size := len(nums)
// 	res := make([]int, size)

// 	for i := 0; i < size; i++ {
// 		res[i] = -1
// 	}

// 	stack := []int{0}

// 	for i := 1; i < 2*size; i++ {
// 		num := nums[i%size]
// 		top := len(stack) - 1

// 		for top >= 0 && nums[stack[top]] < num {
// 			res[stack[top]] = num
// 			stack = stack[:top]
// 			top--
// 		}

// 		if i < size {
// 			stack = append(stack, i)
// 		}
// 	}

// 	return res
// }


// 42. Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/description/
// https://go.dev/play/p/lC590ri9vjw

// func main() {
// 	fmt.Println("Hello, 世界")
// 	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
// 	res := trap(height)
// 	fmt.Println("result================>", res)
// }

// func trap(height []int) int {
// 	lmax, Rmax, total := 0, 0, 0
// 	l, r := 0, len(height)-1

// 	for l < r {
// 		if height[l] <= height[r] {
// 			if lmax > height[l] {
// 				total += lmax - height[l]
// 			} else {
// 				lmax = height[l]
// 			}
// 			l++
// 		} else {
// 			if Rmax > height[r] {
// 				total += Rmax - height[r]
// 			} else {
// 				Rmax = height[r]
// 			}
// 			r--
// 		}

// 		fmt.Println("lamx , rmax", lmax, Rmax)
// 		fmt.Println(total)
// 	}

// 	return total
// }


// 