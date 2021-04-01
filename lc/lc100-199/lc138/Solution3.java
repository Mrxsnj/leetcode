package lc138;

public class Solution3 {

  public Node copyRandomList(Node head) {
    if (head == null) return null;

    Node ptr = head;

    // weave
    while (ptr != null) {
      // Generate
      Node newNode = new Node(ptr.val);
      // Insert
      newNode.next = ptr.next;
      ptr.next = newNode;
      // Move ptr
      ptr = newNode.next;
    }

    // handling random ptr
    ptr = head;
    while (ptr != null) {
      ptr.next.random = (ptr.random == null) ? null : ptr.random.next;
      ptr = ptr.next.next;
    }

    // unweave
    ptr = head;
    Node ptrNewList = head.next;
    Node headOfNewList = head.next;
    while (ptr != null) {
      ptr.next = ptr.next.next;
      ptrNewList.next = (ptrNewList.next == null) ? null : ptrNewList.next.next;
      ptr = ptr.next;
      ptrNewList = ptrNewList.next;
    }

    return headOfNewList;
  }
}
