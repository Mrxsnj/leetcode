从target开始深度优先遍历，除了左右孩子，还要加上父节点（先遍历一遍用哈希表存下来）

递归有3个方向，需要判断下一轮递归中的目标节点是否为来源节点，来避免死循环。

深度加减交由递归去做，将它作为参数。自增自减的话在target节点容易出问题，多+1。
