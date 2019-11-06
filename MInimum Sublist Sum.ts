function minSubListSum(nums: number[], limit: number): number {
  let arr: number[] = new Array(nums.length + 1).fill(0);

  for (let i = 1; i < nums.length + 1; i++) {
    let sumTmp: number = 0;
    let max: number = nums[i - 1];
    arr[i] = arr[i - 1] + nums[i - 1];
    for (let j = i - 1; j >= 0; j--) {
      max = max > nums[j] ? max : nums[j];
      sumTmp += nums[j];
      if (sumTmp > limit) {
        break;
      }
      arr[i] = arr[i] < arr[j] + max ? arr[i] : arr[j] + max;
    }
  }
  return arr[nums.length];
}

console.log(minSubListSum([2, 2, 2, 8, 1, 8, 2, 1], 17));
