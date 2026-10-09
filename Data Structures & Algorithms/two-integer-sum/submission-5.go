import "slices"

func twoSum(nums []int, target int) []int {
    var result []int
    for i := 0; i < len(nums); i++ {
        hash := target - nums[i]
        index := slices.Index(nums, hash)
        if index == i {
            continue
        }
        if index == -1 {
            continue
        } 
        if i < index {
            result = append(result, i, index)
            break
        } else {
            result = append(result, index, i)
            break
        }
           
    }
    return result
}
