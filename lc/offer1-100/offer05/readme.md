# Offer 05. 替换空格

# Q

请实现一个函数，把字符串 `s` 中的每个空格替换成"%20"。

**示例 1：**

```
输入：s = "We are happy."
输出："We%20are%20happy."
```

## Notes

### 算法流程：

1. 初始化一个 list (Python) / StringBuilder (Java) ，记为 res ；
2. 遍历列表 s 中的每个字符 c ：
   - 当 c 为空格时：向 res 后添加字符串 "%20" ；
   - 当 c 不为空格时：向 res 后添加字符 c 
   - 
3. 将列表 res 转化为字符串并返回。
### 复杂度分析：
- 时间复杂度 O(N) ： 遍历使用 O(N) ，每轮添加（修改）字符操作使用 O(1) ；
- 空间复杂度 O(N)： Python 新建的 list 和 Java 新建的 StringBuilder 都使用了线性大小的额外空间。

