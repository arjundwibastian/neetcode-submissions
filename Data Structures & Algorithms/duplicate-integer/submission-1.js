class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        const duplicates = new Set(nums);
        if (nums.length !== duplicates.size) return true;
        else if (nums.length === duplicates.size) return false;
    }
}
