package offer06;

import utils.ListNode;

import java.util.LinkedList;


/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Solution {
  public int[] reversePrint(ListNode head) {
    if (head == null) return new int[]{};

    LinkedList<Integer> list = new LinkedList<>();

    while (head != null) {
      list.addLast(head.val);
      head = head.next;
    }

//    System.out.println(list);
//    System.out.println(list.size());

    int [] retval = new int[list.size()];

    for (int i = 0; i < retval.length; i++) {
      retval[i] = list.removeLast();
    }

    return retval;
  }
}
