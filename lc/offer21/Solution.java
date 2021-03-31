package offer21;

import java.util.Arrays;

public class Solution {
  public int[] exchange(int[] nums) {

    if (nums.length == 0) return nums;

    int left = 0;
    int right = nums.length - 1;

    while (left < right) {
      while (left <= nums.length - 1 && nums[left] % 2 == 1) {
        left++;
      }

      while (right >= 0 && nums[right] % 2 == 0) {
        right--;
      }

      // System.out.println(left);
      // System.out.println(right);

      if (left < right) {
        int tmp = nums[left];
        nums[left] = nums[right];
        nums[right] = tmp;

        // System.out.println(Arrays.toString(nums));
      }
    }

    return nums;
  }
}
