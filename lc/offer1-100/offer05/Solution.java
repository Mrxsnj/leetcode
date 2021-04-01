package offer05;

public class Solution {
  public String replaceSpace(String s) {
    char[] arr = new char[s.length() * 3];
    int newLength = 0;

    for (int i = 0; i < s.length(); i++) {
      char c = s.charAt(i);

      if (c == ' ') {
        arr[newLength++] = '%';
        arr[newLength++] = '2';
        arr[newLength++] = '0';
      } else {
        arr[newLength++] = c;
      }
    }

    return new String(arr, 0, newLength);
  }
}
