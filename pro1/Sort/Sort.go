package Sort

func BubbleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		flag := false
		for j := 1; j <= len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				temp := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = temp
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}
