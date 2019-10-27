/*
 * @lc app=leetcode id=102 lang=javascript
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (50.05%)
 * Likes:    1878
 * Dislikes: 47
 * Total Accepted:    452.1K
 * Total Submissions: 892.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
 * 
 * 
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 
 * return its level order traversal as:
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */

// v1, recursion, dfs
var levelOrder = function (root, level = 0, res = []) {
    if (!root) return res

    if (res[level]) {
        res[level].push(root.val)
    } else {
        res[level] = [root.val]
    }
    levelOrder(root.left, level + 1, res)
    levelOrder(root.right, level + 1, res)
    return res
};

// v2, iteration, bfs
var levelOrder = function (root) {
    if (!root) return []

    const res = []
    const queue = [root]
    while (queue.length) {
        let currentLen = queue.length
        let tmp = []
        for (let i = 0; i < currentLen; i++) {
            const cur = queue.shift()
            cur.left && queue.push(cur.left)
            cur.right && queue.push(cur.right)
            tmp.push(cur.val)
        }
        res.push(tmp)
    }
    return res
}

// @lc code=end