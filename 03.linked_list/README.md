# 链表（Linked List）

## 链表与数组对比

**链表：**

- 不需要连续的存储空间，可以通过指针讲一组零散的内存块串联起来使用
- 对内存消耗比数组高，因为需要额外的存储空间去存储指向下一个节点的指针
- 插入删除时间复杂度：O(1)
- 随机访问时间复杂度：O(n)

**数组：**

- 需连续的存储空间，如果空间不足，需申请更大的内存空间，再把原数组拷贝进去，费时
- 简单易用，可以借助CPU的缓存机制，预读数组中的数据，访问效率更高
- 插入删除时间复杂度：O(n)
- 随机访问时间复杂度：O(1)

## 应用举例

- LRU（Least Recently Used）缓存淘汰算法

---

## 常见单链表问题

### 1.反转链表

时间复杂度：O(n)

### 2.环的检测

典型的算法包括通过hash存储检验和龟兔赛跑。

- **hash存储**

将链表中每个节点的地址存在hash表中，遍历链表，在每个节点处判断是否当前地址已经存在于hash表中，若存在则存在环。

缺点：需要额外的空间存储hash表。

- **龟兔赛跑**

又称`Floyd's Tortoise and Hare`，核心在于一快一慢两个指针，快指针速度是慢指针的两倍，一旦两个指针相遇，则一定存在环。

Stack Overflow中 [一个回答](https://stackoverflow.com/a/54850855/1594792) 下面的图很清晰明了，贴在这里便于理解：

![快慢指针判断环的存在](https://i.stack.imgur.com/rbtDK.png)

简而言之，假设上图中两指针都从起点出发，在p点相遇时：
- 快指针走的距离为：`a + b + c + b = a + 2b + c`
- 慢指针走的距离为：`a + b`
- 由于快指针速度为慢指针的2倍，所以：`a + 2b + c = 2(a + b) => a = c`

因此，即使两个指针不从起点一起出发：慢指针从起点出发，走到q点，此时快指针刚好可以从p点出发，走到q点，二者依然可以在q点相遇。

在[一个博客](https://kchen.cc/2018/11/06/single-circle-linkedlist-entry-point/)中看到了很鲜活的动图，一并放在这里。

![龟兔赛跑测环的动图](http://data.kchen.cc/mac_af-b9a73ed3596db8a22fbc97f6f6abc35f.gif-origin)

### 3.两个有序链表的合并

用两个指针对两个链表从头遍历，分别比较指针指向的节点的大小，小的一方向对应指针后移动。注意处理一方到链尾另一方没有到的情况。

### 4.删除链表倒数第n个节点

可以参考上面检测环的思路，利用一快一慢两个指针，保持两指针的距离为n，当快指针到链尾时，慢指针即为倒数第n个节点，删除即可。

### 5.求链表的中间节点

可以参考上面检测环的思路，利用一快一慢两个指针，快指针的速度是慢指针的2倍，当快指针到链尾时，慢指针指向的就是中间节点。
注意考虑链长为偶数时，中间节点是两个。所以用临时变量保存了循环中未移动之前的slow指针。


## 常见循环链表问题

### 1.约瑟夫问题

## Go 实现

意外发现 Go 的内库中有 list（双向链表）、ring（环）和 heap（堆）的实现。可以学习一下。

## 对应LeetCode题目

解法固定，熟能生巧

- []https://leetcode-cn.com/problems/reverse-linked-list/
- []https://leetcode-cn.com/problems/swap-nodes-in-pairs
- []https://leetcode-cn.com/problems/linked-list-cycle
- []https://leetcode-cn.com/problems/linked-list-cycle-ii
- []https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

## 小结

- 似乎链表的操作中，利用一快一慢两个指针的思路很常见。

# 跳表（Skip List）

为什么要有跳表？我们知道链表prepend、append、insert、delete的时间复杂度都是 O(1)，唯有lookup的时间复杂度是 O(n)，那么如何给链表加速呢？

添加第一级索引、第二级索引，以此类推增加多级索引。

可以学习跳表这种升维度 + 空间换时间的思想。

## 时间复杂度分析

若 n 为节点数量，则第一级索引节点数为 n/2，第二级为 n/4，第三级为 n/8，第 k 级为 n/(2^k)。
假设索引有 h 级，最高级的索引有 2 个节点，n/(2^h) = 2，则 h = log2(n)-1。
因此，在跳表中查询任意数据的时间复杂度为 O(logn)。

## 应用举例

Redis：[https://redisbook.readthedocs.io/en/latest/internal-datastruct/skiplist.html](https://redisbook.readthedocs.io/en/latest/internal-datastruct/skiplist.html)

# Homework

[]https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/ 
[]https://leetcode-cn.com/problems/rotate-array/ 
[]https://leetcode-cn.com/problems/merge-two-sorted-lists/
[]https://leetcode-cn.com/problems/merge-sorted-array/
[]https://leetcode-cn.com/problems/two-sum/
[]https://leetcode-cn.com/problems/move-zeroes/
[]https://leetcode-cn.com/problems/plus-one/