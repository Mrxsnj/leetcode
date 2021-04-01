package lc138;

import java.util.HashMap;
import java.util.Map;

/**
 * iteration
 */

public class Solution2 {

  Map<Node, Node> visited = new HashMap<>();

  public Node copyRandomList(Node head) {
    if (head == null) return null;

    Node ptr = head;

    while (ptr != null) {
      Node newNode = getNode(ptr);
      newNode.next = getNode(ptr.next);
      newNode.random = getNode(ptr.random);

      ptr = ptr.next;
    }

    return visited.get(head);
  }

  private Node getNode(Node node) {
    if (node == null) return null;
    else if (visited.containsKey(node))  return visited.get(node);
    else {
      Node newNode = new Node(node.val);
      visited.put(node, newNode);
      return newNode;
    }
  }
}
