package zdpgo_algorithm

// Jaccard 计算两个数组之间的Jaccard相似度
// 时间复杂度：O(n)
// @param arr1 数组1
// @param arr2 数组2
// @return float64 相似度
func Jaccard[T Number](arr1 []T, arr2 []T) float64 {
	// 边界情况
	if len(arr1) == 0 || len(arr2) == 0 {
		return 0
	}

	// 将两个数组转换为字典
	m1 := make(map[T]struct{}, len(arr1))
	m2 := make(map[T]struct{}, len(arr2))
	for _, v := range arr1 {
		m1[v] = struct{}{}
	}
	for _, v := range arr2 {
		m2[v] = struct{}{}
	}

	// 计算交集的元素个数
	var count float64
	for k, _ := range m1 {
		if _, ok := m2[k]; ok {
			count++
		}
	}

	// 使用算法公式计算相似度
	// 交集个数 / (集合1个数 + 集合2个数 - 交集个数)
	// 由于结果是浮点数类型，需要手动将结果转换为浮点数类型
	return count / float64(len(arr1)+len(arr2)-int(count))
}

// JaccardSorted 用于两个有序数组的快速Jaccard相似度算法
// 时间复杂度：O(n)
// @param arr1 数组1，要求是有序的
// @param arr2 数组2，要求是有序的
// @return float64 相似度
func JaccardSorted[T Number](arr1 []T, arr2 []T) float64 {
	if len(arr1) == 0 || len(arr2) == 0 {
		return 0
	}

	// 求交集的个数
	count := 0
	for i, j := 0, 0; i < len(arr1) && j < len(arr2); {
		// 两个有序的数组，只有其中的某个片段是连续相同的
		if arr1[i] == arr2[j] {
			// 这种情况说明重叠的部分已经出现了
			count++
			i++
			j++
		} else if arr1[i] < arr2[j] {
			// 这种情况说明重叠的部分在arr1的后面，让arr1的索引往后递增
			i++
		} else {
			// 这种情况说明重叠的部分在arr2的后面，让arr2的索引往后递增
			j++
		}
	}

	// 计算相似度
	return float64(count) / float64(len(arr1)+len(arr2)-count)
}
