/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	res := &ListNode{0, nil}
	quick, slow := res, res
	var tmp int
	for p1 != nil && p2 != nil {
		// tmp为进位， Val为值
		quick.Val = (p1.Val + p2.Val + tmp) % 10
		tmp = (p1.Val + p2.Val + tmp) / 10

		// result 创建新结点，依次后传
		quick.Next = &ListNode{0, nil}
		slow = quick
		quick, p1, p2 = quick.Next, p1.Next, p2.Next
	}

	if p1 == nil && p2 == nil {
		quick.Val = tmp
	}

	for p1 != nil {
		quick.Val = (p1.Val + tmp) % 10
		tmp = (p1.Val + tmp) / 10

		quick.Next = &ListNode{0, nil}
		slow = quick
		quick, p1 = quick.Next, p1.Next
	}

	for p2 != nil {
		quick.Val = (p2.Val + tmp) % 10
		tmp = (p2.Val + tmp) / 10

		quick.Next = &ListNode{0, nil}
		slow = quick
		quick, p2 = quick.Next, p2.Next
	}

	if quick.Val == 0 && tmp == 0 {
		slow.Next = nil
	} else {
		quick.Val = tmp
	}

	return res
}

// 快慢指针，使用慢指针的时候注意改变的是Next，改变自身毫无意义