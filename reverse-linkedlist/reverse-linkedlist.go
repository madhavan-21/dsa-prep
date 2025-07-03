//Reversing a linked list


// Reverse a linked list
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}

	return prev
}


//example usage
//1.https://leetcode.com/problems/reverse-linked-list/
//2.https://leetcode.com/problems/reverse-linked-list-ii/submissions/1684052023/