
var totalNQueens = function(n) {
  var all = Math.pow(2, n) - 1;
  var solutionCount = 0;

  var findSolutions = function(cols, ld, rd) {
    var pos = ~(cols | ld | rd) & all;
    while (pos > 0) {
      var bit = -pos & pos;
      pos = pos ^ bit;
      let newCol = cols | bit;
      let newLD = (ld | bit) << 1;
      let newRD = (rd | bit) >> 1;
      findSolutions(newCol, newLD, newRD);
    }
    if (cols === all) {
      solutionCount++;
    }
  };
  findSolutions(0, 0, 0);
  return solutionCount;
};

const n = 14
console.log(totalNQueens(n))
