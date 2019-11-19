/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 *
 * https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (51.32%)
 * Likes:    1071
 * Dislikes: 438
 * Total Accepted:    288.1K
 * Total Submissions: 561K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers that is already sorted in ascending order, find
 * two numbers such that they add up to a specific target number.
 * 
 * The function twoSum should return indices of the two numbers such that they
 * add up to the target, where index1 must be less than index2.
 * 
 * Note:
 * 
 * 
 * Your returned answers (both index1 and index2) are not zero-based.
 * You may assume that each input would have exactly one solution and you may
 * not use the same element twice.
 * 
 * 
 * Example:
 * 
 * 
 * Input: numbers = [2,7,11,15], target = 9
 * Output: [1,2]
 * Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 * 
 */
// ![two-pointer]
func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers) - 1
    for i<j {
        actual := numbers[i] + numbers[j]
        if actual == target {
            return []int{i+1, j+1}
        } else if actual < target {
            i++
        } else {
            j--
        }                    
    }
    return []int{}
}

