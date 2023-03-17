package utils

func IntSum(a int, b int) int {
	return a + b
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 设置一个标志变量，记录是否发生了交换
		swapped := false
		// 每次循环后，最大的元素会沉到数组末尾，所以可以减少比较次数
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换相邻的两个元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 标记发生了交换
				swapped = true
			}
		}
		// 如果没有发生交换，说明已经有序，可以提前结束循环
		if !swapped {
			break
		}
	}
}
