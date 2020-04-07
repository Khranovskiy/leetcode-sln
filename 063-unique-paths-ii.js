class Matrix1D {
  constructor(m, n) {
    this.m = m
    this.n = n
    this.storage = new Array(m * n).fill(0)
    // let mat = [...Array(m)].map(item => Array(n).fill(0))
  }
  get(r, c) {
    return this.storage[r * this.n + c]
  }
  set(r, c, value) {
    return (this.storage[r * this.n + c] = value)
  }
}
/**
 * @param {number[][]} obstacleGrid
 * @return {number}
 */
var uniquePathsWithObstacles2 = function(obstacleGrid) {
  let m = obstacleGrid.length
  let n = obstacleGrid[0].length
  if (m === n && n === 1) {
    return 1 - obstacleGrid[0][0]
  }
  // [0,0,0],
  // [0,1,0],
  // [0,0,0]

  // 02 01 01
  // 01 XX 01
  // 01 01 01
  let mat = new Matrix1D(m, n)
  const calcAddition = (isLast, rNext, cNext) => {
    let addition = 0
    if (!isLast) {
      let noObstacle = obstacleGrid[rNext][cNext] === 0
      if (noObstacle) {
        addition = mat.get(rNext, cNext)
      }
    }
    return addition
  }
  const lastRow = m - 1
  const lastColumn = n - 1
  for (let r = lastRow; r >= 0; r--) {
    for (let c = lastColumn; c >= 0; c--) {
      if (obstacleGrid[r][c] !== 0) {
        continue
      }
      let value = 0
      if (r === lastRow && c === lastColumn) {
        value = 1
      } else {
        value += calcAddition(r === lastRow, r + 1, c)
        value += calcAddition(c === lastColumn, r, c + 1)
      }
      mat.set(r, c, value)
    }
  }
  return mat.get(0, 0)
}

var uniquePathsWithObstacles = function(obstacleGrid) {
  const m = obstacleGrid.length
  if (m === 0) return 0
  const n = obstacleGrid[0].length

  let mat = new Matrix1D(m + 1, n + 1)
  mat.set(m - 1, n, 1)

  for (let r = m - 1; r >= 0; r--) {
    for (let c = n - 1; c >= 0; c--) {
      let value = obstacleGrid[r][c] === 1 ? 0 : mat.get(r, c + 1) + mat.get(r + 1, c)
      mat.set(r, c, value)
    }
  }
  return mat.get(0, 0)
}
