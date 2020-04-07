function maxSubArrayK(nums) {
  const n = nums.length
  let max_cur = nums[0]
  let max_global = max_cur
  for (let cur, ii = 1; ii < n; ++ii) {
    cur = nums[ii]
    max_cur = Math.max(cur, max_cur + cur)
    max_global = Math.max(max_global, max_cur)
  }
  return max_global
}

//O(n log n) runtime, O(log n) stack space – Divide and Conquer
function maxSubArray(nums) {
  return maxSubArrayDC(nums, 0, nums.length - 1)
}
//[2, 1, –3, 4, –1, 2, 1, –5, 4]
function maxSubArrayDC(nums, left, right){
  // console.log(`left - right ${left} ${right}`)
  if(left > right){
    return Number.MIN_SAFE_INTEGER
  }
  const middle = left + Math.trunc((right - left) / 2)
  let leftAnswer = maxSubArrayDC(nums, left, middle - 1)
  let rightAnswer = maxSubArrayDC(nums, middle + 1, right)
  
  let leftMaxSum = 0
  let sum = 0
  for(let ii = middle - 1; ii >= left; ii--){
    sum += nums[ii]
    leftMaxSum = Math.max(leftMaxSum, sum)
  }
  let rightMaxSum = 0
  sum = 0
  for(let ii = middle + 1; ii <= right; ii++){
    sum += nums[ii]
    rightMaxSum = Math.max(rightMaxSum, sum)
  }
  let a = Math.max(leftAnswer, rightAnswer)
  let b = leftMaxSum + nums[middle] + rightMaxSum
  return Math.max(a, b)
}