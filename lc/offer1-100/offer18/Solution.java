package offer18;

import utils.ListNode;

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */

public class Solution {
  public ListNode deleteNode(ListNode head, int val) {

    if (head.val == val) {
      return head.next;
    }

    ListNode prev = head;
    ListNode curr = head.next;

    while (curr != null) {
      if (curr.val == val) {
        prev.next = curr.next;
        curr = curr.next;
        break;
      }

      curr = curr.next;
      prev = prev.next;
    }

    return head;
  }
}
