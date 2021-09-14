/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
  var dic = new Set()
  const n = nums.length
  for(let curr, ii = 0; ii < n; ++ii){
    curr = nums[ii]
    if (dic.has(curr)){
      dic.delete(curr)
    }else{
      dic.add(curr)
    }
  }
  return Array.from(dic)[0]
};
