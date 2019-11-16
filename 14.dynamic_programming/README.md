# 动态规划

## 关键点

- 动态规划 和 递归 或者 分治 没有根本上的区别（关键看**有无最优的子结构**，参[120_test.go](./test/120_test.go)中的错误解法）
- **共性：找到重复子问题**
- 差异性：最优子结构、中途可以**淘汰**次优解。三步走：
    1. 最优子结构 opt[n] = best_of(opt[n-1], opt[n-2], ...)
    2. 储存中间状态：opt[i]
    3. 递推公式（美其名曰：状态转移方程 或 DP 方程）
        - Fib：opt[i] = opt[i-1] + opt[i-2]
        - 二维路径：opt[i, j] = opt[i+1, j] + opt[i, j+1]（且判断 opt[i, j]是否为空地）

## 斐波那契数列（一维数组）

递推公式：
```
F(n) = F(n-1) + F(n-2)
```

1. 记忆化搜索（memozation）：加缓存，自顶向下
2. 自底向上：递推（for 循环，所以又叫动态递推），eg: 0，1，1，2，3，5，8，13，……

## 路径计数

可以转换思路，变为自底向上，从右下开始往左上走。

递推公式：（opt = optimal）
```
opt[i, j] = opt[i+1, j] + opt[i, j+1]
```

完整逻辑：
```
if a[i, j] = '空地':
    opt[i, j] = opt[i+1, j] + opt[i, j+1]
else:
    opt[i, j] = 0   // 遇到障碍物，走法为 0
```

## 小结

1. 打破自己的思维惯性，形成机器思维（找重复性）
2. 理解复杂逻辑的关键
3. 也是职业进阶的要点要领

## 备注

对于 1143题 LCS 问题（Longest Common Subsequence），个人觉得如下图示更容易理解：

![1143lcs](./1143lcs.png)
(原图来自：[https://leetcode.com/problems/longest-common-subsequence/discuss/348884/C%2B%2B-with-picture-O(nm)](https://leetcode.com/problems/longest-common-subsequence/discuss/348884/C%2B%2B-with-picture-O(nm)))

即在原有 `m * n` 表格的基础上，最前面增加一行一列，变成 `(m+1) * (n+1)` 的表格，这样在迭代求解时 for 循环的范围就变成了 `[1, m]`，方便比较上一行/上一列与当前行/列的关系，并更新数据，最终结果位于 `dp[m][n]`。