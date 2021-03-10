package lc739;

import java.util.Arrays;

public class Solution {

  /***
   * Decreasing stack saves the order of the temperatures(not the value).
   * @param T e.g., [73,74,75,71,69,72,76,73]
   * @return retval e.g., [1,1,4,2,1,1,0,0]
   */
  public int[] dailyTemperatures(int[] T) {
    int[] retval = new int[T.length];
    int[] stack = new int[T.length];
    // init with 0
    // System.out.println(Arrays.toString(stack));
    int i = -1; // stack ptr

    for (int j = 0; j < T.length; j++) {
      // stack not null
      while (i != -1 && T[j] > T[stack[i]]) {
        retval[stack[i]] = j - stack[i];
        i--;
      }

      stack[++i] = j;
    }

    return retval;
  }
}
