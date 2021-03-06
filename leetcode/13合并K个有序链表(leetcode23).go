package leetcode

/*
归并思想，2个一组，两两合并
*/
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeList(lists, 0, len(lists)-1)
}

func mergeList(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	return mergeTwoLists(mergeList(lists, left, mid), mergeList(lists, mid+1, right))
}
