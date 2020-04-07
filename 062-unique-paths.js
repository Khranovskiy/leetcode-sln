/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
var uniquePathsBacktrack = function(m, n) {
  return backtrackMemo(0, 0, m, n)
}
//Backtracking Solution using Memoization
var backtrackMemo = function(r, c, m, n, memo) {
  memo = memo || new Array(101 * 100) // {}
  // if (memo[n]) {
  //     return memo[n]
  // }
  if (r === m - 1 && c === n - 1) return 1
  if (r >= m || c >= n) return 0

  const keyBottom = getKey(r + 1, c)
  const keyRight = getKey(r, c + 1)
  if (memo[keyBottom] === undefined) {
    memo[keyBottom] = backtrackMemo(r + 1, c, m, n, memo)
  }
  if (memo[keyRight] === undefined) {
    memo[keyRight] = backtrackMemo(r, c + 1, m, n, memo)
  }
  // return memo[n] = fib..
  return memo[keyBottom] + memo[keyRight]
}
var getKey = function(r, c) {
  return r * 101 + c
}

// O(mn) runtime, O(mn) space – Bottom-up dynamic programming

// The total unique paths at grid (r, c) are equal to the sum
// of total unique paths from grid to the right (r, c + 1) and
// the grid below (r + 1, c).

// We observe that all grids of the bottom edge and right edge
// must all have only one unique path to the bottom-right
// corner. Using this as the base case, we can build our way
// up to our solution at grid (1, 1) using the relationship above.

// 28 21 15 10 06 03 01
// 07 06 05 04 03 02 01
// 01 01 01 01 01 01 01
var uniquePaths2 = function(m, n) {
  // let mat = new Array(m).fill(new Array(n).fill(1))
  let mat = [...Array(m)].map(item => Array(n).fill(1))
  for (let r = m - 1 - 1; r >= 0; r--) {
    for (let c = n - 1 - 1; c >= 0; c--) {
      mat[r][c] = mat[r + 1][c] + mat[r][c + 1]
    }
  }
  return mat[0][0]
}
// 3 x 2:
// 03 01
// 02 01
// 01 01

// 3 x 3:
// 06 03 01
// 03 02 01
// 01 01 01

// 2 x 2:
// 02 01
// 01 01

class Matrix1D {
  constructor(m, n) {
    this.m = m
    this.n = n
    this.storage = new Array(m * n).fill(1)
  }
  get(r, c) {
    return this.storage[r * this.n + c]
  }
  set(r, c, value) {
    return (this.storage[r * this.n + c] = value)
  }
}

var uniquePaths3 = function(m, n) {
  let mat = new Matrix1D(m, n)
  for (let r = m - 1 - 1; r >= 0; r--) {
    for (let c = n - 1 - 1; c >= 0; c--) {
      let value = mat.get(r + 1, c) + mat.get(r, c + 1)
      mat.set(r, c, value)
    }
  }
  return mat.get(0, 0)
}

// Combinatorial Solution:
// 28 21 15 10 06 03 01
// 07 06 05 04 03 02 01
// 01 01 01 01 01 01 01
// Look at the 7×3 sample grid in the picture above.
// Notice that no matter how you traverse the grids, you always traverse a total of 8 steps.
// To be more exact, you always have to choose 6 steps to the right (R)
// and 2 steps to the bottom (B).
//
// Therefore, the problem can be transformed to a question of
// how many ways can you choose 6R‘s and 2B‘s in these 8 steps.
// The answer is ( 8 by 2) or (8 by 6) => the general solution is
// for an m x n grid is
// ( m+n-2 by m - 1)
//
// Runtime: 48 ms, faster than 95.59% of JavaScript online submissions for Unique Paths.
// Memory Usage: 34 MB, less than 54.55% of JavaScript online submissions for Unique Paths.
var uniquePaths = function(m, n) {
  let nn = m + n - 2
  let rr = m - 1
  var memo = []
  return factorial(nn, memo) / (factorial(rr, memo) * factorial(nn - rr, memo))
}
var factorial = function(n, memo) {
  if (n == 0 || n == 1) return 1
  if (memo[n] !== undefined) return memo[n]
  return (memo[n] = factorial(n - 1, memo) * n)
}
