package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 先让一个指针走k个结点，然后两个指针一起走，先走的到末尾的时候，第二个就是倒数k个节点
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	ptr1, ptr2 := head, head
	for i := 1; i < k; i++ {
		ptr1 = ptr1.Next
	}

	for ptr1.Next != nil {
		ptr2 = ptr2.Next
		ptr1 = ptr1.Next
	}
	return ptr2
}
