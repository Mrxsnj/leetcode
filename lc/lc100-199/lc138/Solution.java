package lc138;

import java.util.HashMap;
import java.util.Map;

/**
 * recursion
 */

public class Solution {
  Map<Node, Node> visited = new HashMap<>();

  public Node copyRandomList(Node head) {
    if (head == null) return null;

    if (visited.containsKey(head))  return visited.get(head);

    Node newNode = new Node(head.val);
    visited.put(head, newNode);

    newNode.next = copyRandomList(head.next);
    newNode.random = copyRandomList(head.random);

    return newNode;
  }
}
