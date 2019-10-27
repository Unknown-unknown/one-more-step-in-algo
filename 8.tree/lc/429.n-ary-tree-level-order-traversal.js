/*
 * @lc app=leetcode id=429 lang=javascript
 *
 * [429] N-ary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/n-ary-tree-level-order-traversal/description/
 *
 * algorithms
 * Easy (60.66%)
 * Likes:    349
 * Dislikes: 36
 * Total Accepted:    49K
 * Total Submissions: 80.6K
 * Testcase Example:  '{"$id":"1","children":[{"$id":"2","children":[{"$id":"5","children":[],"val":5},{"$id":"6","children":[],"val":6}],"val":3},{"$id":"3","children":[],"val":2},{"$id":"4","children":[],"val":4}],"val":1}'
 *
 * Given an n-ary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
 * 
 * For example, given a 3-ary tree:
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * We should return its level order traversal:
 * 
 * 
 * [
 * ⁠    [1],
 * ⁠    [3,2,4],
 * ⁠    [5,6]
 * ]
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The depth of the tree is at most 1000.
 * The total number of nodes is at most 5000.
 * 
 * 
 */

// @lc code=start
/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */
/**
 * @param {Node} root
 * @return {number[][]}
 */

// v1, bfs
var levelOrder = function (root) {
    if (!root == null) return []

    const res = []
    const queue = [root]
    while (queue.length > 0) {
        let currentLen = queue.length
        let tmp = []
        for (let i = 0; i < currentLen; i++) {
            let node = queue.shift()
            console.log("⬆️  shift from queue:", node.val, queue)
            if (node.children) {
                for (const child of node.children) {
                    queue.push(child)
                    console.log("⬇️  push to queue:", child.val, queue)
                }
            }
            tmp.push(node.val)
            console.log("⏬  push to tmp:", node.val, tmp)
        }
        res.push(tmp)
    }

    return res
};

// for test
// function Node(val, children) {
//     this.val = val;
//     this.children = children;
// };

// let children = [new Node(1, [new Node(4, null), new Node(5, null)]), new Node(2, [new Node(6, null)]), new Node(3, null)]
// let root = new Node(0, children)
// let res = levelOrder(root)
// console.log(res)
// @lc code=end