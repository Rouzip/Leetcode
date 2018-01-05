/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function(head) {
	let result = head;
	if (!head || !head.next){
		return head;
	}
	while(head.next){
		if (head.val == head.next.val){
			head = head.next;
			continue;
		}
		result.next = head;
	}
	return result;
};