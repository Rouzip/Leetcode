/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* rotateRight(ListNode* head, int k) {
        if (head == NULL || head->next == NULL || k == 0){
            return head;
        }
        int length = 1;
        ListNode* cur = head;
        ListNode* newH;
        while(cur->next) {
            cur = cur->next;
            length++;
        }
        cur->next = head;
        k = length - k%length;
        while(k--) {
            cur = cur->next;
        }
        newH = cur->next;
        cur->next = NULL;
        return newH;
    }
};

/**
 * 这道题使用了循环链表来进行寻找，避免了使用多个指针来进行判断，同时减少了边界问题的判断
 * 判断边界的时候需要谨慎，很可能超出边界。（好久没用指针我都忘了语法了= =
 */