package v2024

// SearchBinary 二分查找算法的基本实现
// @param arr 要查找数组，必须是有序的
// @param target 要查找的目标值
// @return int 返回的索引，找到了返回具体的索引，找不到返回-1
func SearchBinary[T Number](arr []T, target T) int {
	// 计算开始索引和结束索引
	begin := 0
	end := len(arr) - 1

	// 遍历所有的可能性，在开始索引小于或等于结束索引的时候一直找
	for begin <= end {
		// 计算中间索引
		mid := (begin + end) / 2 // go 语言里面的除法默认就是整除

		// 找到了
		if arr[mid] == target {
			return mid
		}

		// 中间值比目标值大，去左边找
		if arr[mid] > target {
			end = mid - 1
		} else {
			// 中间值比目标值小，去右边找
			begin = mid + 1
		}
	}

	// 循环完了都没有找到
	return -1
}
