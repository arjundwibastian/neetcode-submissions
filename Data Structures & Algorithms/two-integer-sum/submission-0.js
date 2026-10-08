class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        let contain = [];
        nums.forEach(function (value, index) {
            nums.forEach( function (val, i) {
                if (index !== i && contain.length === 0) {
                    const sum = value + val;
                    if (sum === target) contain.push(index, i)
                }
            })
        })
        return contain;
    }
}
