# Offer58. 左旋转字符串

字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

**示例 1：**

```
输入: s = "abcdefg", k = 2
输出: "cdefgab"
```

**示例 2：**

```
输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"
```

## Notes

### 方法1 切片

go用s[n:]

Java用s.substring()

### 方法2 列表遍历拼接

```java
StringBuilder res = new StringBuilder();
for(int i = n; i < s.length(); i++)
    res.append(s.charAt(i));
for(int i = 0; i < n; i++)
    res.append(s.charAt(i));
return res.toString();
```

### 方法3 字符串遍历拼接

```java
String res = "";
for(int i = n; i < s.length(); i++)
    res += s.charAt(i);
for(int i = 0; i < n; i++)
    res += s.charAt(i);
return res;
```

