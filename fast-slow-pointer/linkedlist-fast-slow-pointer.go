//Linked list: fast and slow pointer
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func fn(head *ListNode) int {
    slow, fast := head, head
    ans := 0

    for fast != nil && fast.Next != nil {
        // do logic
        slow = slow.Next
        fast = fast.Next.Next
    }

    return ans
}


//https://leetcode.com/problems/middle-of-the-linked-list/


// func middleNode(head *ListNode) *ListNode {
//     slow , fast := head , head

//     for fast !=nil && fast.Next !=nil{
//         slow = slow.Next
//         fast = fast.Next.Next
//     }
//     return slow
// }


//https://leetcode.com/problems/linked-list-cycle/
// func hasCycle(head *ListNode) bool {
// 	slow, fast := head, head

// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next

// 		if slow == fast {
// 			return true
// 		}
// 	}

// 	return false
// }



//https://leetcode.com/problems/linked-list-cycle-ii/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func detectCycle(head *ListNode) *ListNode {
// 	slow, fast := head, head

// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 		if slow == fast {
// 			break
// 		}
// 	}

// 	if fast == nil || fast.Next == nil {
// 		return nil
// 	}

// 	slow = head

// 	for slow != fast {
// 		slow = slow.Next
// 		fast = fast.Next
// 	}
// 	return slow

// }

