package lc155;

import java.util.LinkedList;

public class MinStack {
  LinkedList<Integer> stack;
  LinkedList<Integer> minStack;

  /**
   * initialize your data structure here.
   */
  public MinStack() {
    stack = new LinkedList<>();
    minStack = new LinkedList<>();
    minStack.push(Integer.MAX_VALUE);
  }

  public void push(int val) {
    stack.addLast(val);

    int prev = minStack.removeLast();
    minStack.addLast(prev);
    minStack.addLast(prev < val ? prev : val);
  }

  public void pop() {
    stack.removeLast();
    minStack.removeLast();
  }

  public int top() {
    int retval = stack.removeLast();
    stack.addLast(retval);
    return retval;
  }

  public int getMin() {
    int retval = minStack.removeLast();
    minStack.addLast(retval);
    return retval;
  }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(val);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */
