package lc

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 双索引遍历法，时间复杂度O(2n)
	o := 0
	h := head
	pre := &ListNode{Next: head}
	preItor := pre
	for h != nil {
		h = h.Next
		if o < n {
			o++
			continue
		}
		preItor = preItor.Next
	}
	deleteNode := preItor.Next
	preItor.Next = deleteNode.Next
	deleteNode.Next = nil
	return pre.Next
}

// @lc code=end
