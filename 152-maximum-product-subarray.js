// https://leetcode.com/problems/maximum-product-subarray/submissions/
function maxProduct(nums) {
  const n = nums.length
  let max = nums[0],
    min = nums[0],
    max_answer = max
  for (let ii = 1; ii < n; ++ii) {
    let cur = nums[ii], maxTmp = max, minTmp = min
    max = Math.max(Math.max(cur, maxTmp * cur), minTmp * cur)
    min = Math.min(Math.min(cur, maxTmp * cur), minTmp * cur)
    max_answer = Math.max(max_answer, max)
  }
  return max_answer
}
//Runtime: 48 ms, faster than 98.37% of JavaScript online submissions for Maximum Product Subarray.
maxProduct([2, 3, -2, 4])
