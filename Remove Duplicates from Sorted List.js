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
	for (let node = head;node;node = node.next){
		while(node.next && node.val == node.next.val)
			node.next = node.next.next;
	}
  return head;
};
//链表去重，对于链表的操作还需要熟练，一开始纠结于各种情况，属于钻牛脚尖了