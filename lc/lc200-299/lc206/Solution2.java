package lc206;

import utils.ListNode;

/**
 * recursion
 */
public class Solution2 {
  public ListNode reverseList(ListNode head) {
    return recur(head, null);
  }

  private ListNode recur(ListNode curr, ListNode prev) {
    if (curr == null) return prev;
    ListNode retval = recur(curr.next, curr);
    curr.next = prev;
    return retval;
  }
}


/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */