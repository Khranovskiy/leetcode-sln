/**
 * @param {number} n
 * @return {number}
 */
var climbStairs2 = function(n) {
  // base cases
  if (n <= 0) return 0
  if (n == 1) return 1
  if (n == 2) return 2

  // The key intuition to solve the problem is that given a number of stairs n,
  // if we know the number ways to get to the points [n-1] and [n-2] respectively,
  // denoted as n1 and n2 , then the total ways to get to the point [n] is n1 + n2.
  // Because from the [n-1] point, we can take one single step to reach [n].
  // And from the [n-2] point, we could take two steps to get there.

  // The solutions calculated by the above approach are complete and non-redundant.
  // The two solution sets (n1 and n2) cover all the possible cases on how the final
  // step is taken. And there would be NO overlapping among the final solutions
  // constructed from these two solution sets, because they differ in the final step.
  let one_step_before = 2 //refers to the number of solutions until the point [n-1]
  let two_steps_before = 1 //refers to the number of solution until the point [n-2]
  let all_ways = 0 // corresponds to the number of solutions to get to the point [n]

  for (let i = 2; i < n; i++) {
    // From the point [n-1], we take one step to reach the point [n].
    // From the point [n-2], we take a two-steps leap to reach the point [n].
    // So it goes without saying that the total number of solution to reach the point [n] should be [n-1] + [n-2].
    all_ways = one_step_before + two_steps_before
    two_steps_before = one_step_before
    one_step_before = all_ways
  }
  return all_ways
}

var climbStairs = function(n) {
  let dp = new Array(n + 1).fill(0)
  if (n == 1) {
    return 1
  }
  if (n == 2) {
    return 2
  }
  dp[0] = 0
  dp[1] = 1
  dp[2] = 2
  for (let i = 3; i <= n; i++) {
    dp[i] = dp[i - 1] + dp[i - 2]
  }
  return dp[n]
}
