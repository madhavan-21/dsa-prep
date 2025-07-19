//Monotonic increasing stack

func monotonicIncreasingStack(arr []int) int {
	stack := []int{}
	ans := 0

	for _, num := range arr {
		// For monotonic decreasing stack: change > to <
		for len(stack) > 0 && stack[len(stack)-1] > num {
			// Do some logic when popping
			stack = stack[:len(stack)-1] // pop
		}

		stack = append(stack, num) // push
	}

	return ans
}

//EXAMPLE: Largest Rectangle in Histogram
//https://leetcode.com/problems/largest-rectangle-in-histogram/
// func largestRectangleArea(heights []int) int {
//     nsl:=make([]int,len(heights))
//     stk:=[]int{}
//     for i:=0;i<len(heights);i++{
//         for len(stk)>0 && heights[stk[len(stk)-1]]>=heights[i]{
//             stk=stk[:len(stk)-1]
//         }
//         if len(stk)==0{
//             nsl[i]=-1
//         }else{
//             nsl[i]=stk[len(stk)-1]
//         }
//         stk=append(stk,i)
//     }

//     nsr := make([]int, len(heights))
// 	stk = []int{} // Reset stack
// 	for i := len(heights) - 1; i >= 0; i-- {
// 		// Pop from the stack until we find a smaller element
// 		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
// 			stk = stk[:len(stk)-1]
// 		}
// 		if len(stk) == 0 {
// 			nsr[i] = len(heights) // No smaller element on the right
// 		} else {
// 			nsr[i] = stk[len(stk)-1] // Index of the last smaller element
// 		}
// 		stk = append(stk, i)
//     }
//     maxArea := 0
// 	for i := 0; i < len(heights); i++ {
// 		width := nsr[i] - nsl[i] - 1
// 		area := heights[i] * width
// 		maxArea = int(math.Max(float64(maxArea), float64(area)))
// 	}

// 	return maxArea
// }