# 树

Linked List 是特殊化的 Tree；Tree 是特殊化的 Graph。

## 二叉树遍历

1. 前序（Pre-order）：根-左-右
2. 中序（In-order）：左-根-右
3. 后序（POST-order）：左-右-根

（所谓前中后，其实是根的位置）

## 二叉搜索树

又称有序二叉树（Ordered Binary Tree）、排序二叉树（Sorted Binary Tree），是指一棵空树或者具有下列性质的二叉树：
1. 左子树上**所有结点**的值均小于它的根结点的值；
2. 右子树上**所有结点**的值均大于它的根结点的值；
3. 以此类推：左、右子树也分别为二叉查找树（这就是**重复性！**）

中序遍历：升序排列

## 二叉搜索树常见操作

1. 查询
2. 插入新结点（创建）
3. 删除

Demo：[https://visualgo.net/zh/bst](https://visualgo.net/zh/bst)

