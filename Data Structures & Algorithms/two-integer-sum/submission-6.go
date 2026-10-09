func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, num := range nums {
        complement := target - num

        if prevIndex, found := seen[complement]; found {
            return []int{prevIndex, i}
        }

        seen[num] = i
    }

    return nil
}
