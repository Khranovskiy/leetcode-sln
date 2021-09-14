var moveZeroes = function(nums) {
  let placement = 0
	for (let i = 0; i < nums.length; i++) {
		if ( nums[i] !== 0) {
			let hold = nums[i]
			nums[i] = nums[placement]
			nums[placement] = hold
			placement++
		}
	}
}
