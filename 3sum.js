//[-1,0,1,2,-1,-4] -> [[-1,-1,2],[-1,0,1]]
function threeSum(nums) {
  let result = []
  nums.sort((a,b)=>a-b)
  
  for(let firstIndex = 0; firstIndex < nums.length - 1 - 1; firstIndex++){
    if(nums[firstIndex] === nums[firstIndex + 1] && nums[firstIndex] !== 0){
      continue
    }
    const target = -nums[firstIndex]
    const hashmap = new Map()
    for (let secondIndex = 0; secondIndex < nums.length; secondIndex++){
      if(secondIndex === firstIndex){
        continue;
      }
      let second = nums[secondIndex]
      if(hashmap.has(target-second)) {
        let thirdIndex = hashmap.get(target-second)
        result.push([nums[firstIndex],nums[secondIndex],nums[thirdIndex] ])
      }
      hashmap.set(second, secondIndex)
    }
  }
  return result
}
// https://www.youtube.com/watch?v=jXZDUdHRbhY