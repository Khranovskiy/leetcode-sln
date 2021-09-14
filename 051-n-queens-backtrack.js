const queenPositionToBoardRow = columnWithQueen => {
  let result = ''
  for (let i = 0; i < n; i++) {
    result += i === columnWithQueen ? 'Q' : '.'
  }
  return result
}

class NQueensBoard {
  constructor(size) {
    this.n = size
    this.board = new Array(size).fill(0)
    this.colsUsed = new Array(size).fill(false)
    this.diago135Used = Array(size * 2 - 1).fill(false)
    this.diago45Used = Array(size * 2 - 1).fill(false)
    this.cols = 0
    this.ld = 0
    this.rd = 0
  }
  make_move(row, column) {
    const i135 = row + this.n - column - 1
    const i45 = row + column

    const oldColumnInRow = this.board[row]
    this.board[row] = column
    this.colsUsed[column] = true
    this.diago135Used[i135] = true
    this.diago45Used[i45] = true

    return oldColumnInRow
  }
  unmake_move(row, column, oldRowState) {
    const i135 = row + this.n - column - 1
    const i45 = row + column

    this.board[row] = oldRowState
    this.colsUsed[column] = false
    this.diago135Used[i135] = false
    this.diago45Used[i45] = false
  }

  *getCandidates(candidateRow) {
    for (let column = 0; column < this.n; column++) {
      const indexOfDiago135 = candidateRow + this.n - column - 1
      const indexOfDiago45 = candidateRow + column

      const badMove =
        this.colsUsed[column] ||
        this.diago135Used[indexOfDiago135] ||
        this.diago45Used[indexOfDiago45]

      if (!badMove) {
        yield column
      }
    }
  }
  isLastRow(row) {
    return row === this.n - 1
  }
}

class Solution {
  constructor() {
    this.result = []
  }
  process(board) {
    this.result.push([...board.board])
  }
  get length() {
    return this.result.length
  }
  // return result.map(board => board.map(r => queenPositionToBoardRow(r)))

  solveNQueens(n) {
    let board = new NQueensBoard(n)
    const initialRow = 0
    for (let column = 0; column < n; column++) {
      backtrack(board, this, initialRow, column)
    }
    return solution.length
  }
}

function backtrack(board, solution, row, col) {
  let old = board.make_move(row, col)
  if (board.isLastRow(row)) {
    solution.process(board)
  } else {
    let nextRow = row + 1
    for (let candidateColumn of board.getCandidates(nextRow)) {
      backtrack(board, solution, nextRow, candidateColumn)
    }
  }
  board.unmake_move(row, col, old)
}

let solution = new Solution();
let result = solution.solveNQueens(11)
process.stdout.write(result.toString())

