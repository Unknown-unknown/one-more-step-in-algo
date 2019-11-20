# AVL树和红黑树

## AVL树

1. 平衡二叉搜索树
2. 每个结点存平衡因子 = 左子树高度 - 右子树高度，取值范围为 `[-1, 0, 1]`
3. 四种旋转操作

**不足：**结点需要存储额外信息、且调整次数频繁

### 旋转操作

- 左左子树 -> 右旋
- 右右子树 -> 左旋
- 左右子树 -> 左右旋
- 右左子树 -> 右左旋


## 红黑树

红黑树是一种**近似平衡**二叉树，它能够确保任何一个结点的左右子树的**高度差小于两倍**。具体来说，红黑树是满足如下条件的二叉搜索树：

- 每个结点要么是红色，要么是黑色
- 根节点是黑色
- 每个叶子结点（NIL结点、空结点）是黑色的
- 不能有相邻接的两个红色结点
- 从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点

### 应用场景

- 在nodejs中，调用`setTimeout()`或者`setInterval()`创建的定时器会被插入到定时器观察者内部的一个红黑树中。每次Tick执行时，会从该红黑树中迭代取出定时器对象，检查是否超过定时时间，如果超过，就形成一个事件，它的回调函数将立即执行。

### 解决的问题

由于红黑树是一种性能非常稳定的二叉查找树，所以在工程中，但凡是用到动态插入、删除、查找数据的场景，都可以用到它。但由于其实现复杂，自己写代码时更倾向用跳表来代替它。

### 下一步思考

动态数据结构支持动态地数据插入、删除、查找操作，除了红黑树，还有哪些？对比各自的优劣 及应用场景？

- 动态数据结构是支持动态的更新操作，里面存储的数据是时刻变化的，也就是说，还支持删除、插入数据。而且，这些操作都很高效。因此，红黑树算是一个 **高效**的动态数据结构。链表、队列、栈实际上算不上，因为操作有限、且查询效率不高。

## AVL vs 红黑树

- AVL trees provide **faster lookups** than Red Black Trees because they are **more strickly balanced**.
- Red Black Trees provide **faster insertion and removal** operations than AVL trees as fewer rotations are done due to relatively relaxed balancing.
- AVL trees store balance **factors or heights** with each node, thus requires storage for an integer per node whereas Red Black Trees requires only 1 bit of information per node.
- Red Black Trees are used in most of the **language libraries like map, multimap, multiset in C++** whereas AVL trees are used in **databases** where faster retrevals are required.