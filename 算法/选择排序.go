package main
/**
* 从未排序序列中找到最小元素存放到起始位置，然后继续从未排序序列中查找最小元素放入已排序序列末尾,直到所有序列排序完成
* 时间复杂度：O(n②) 空间复杂度：O(1)
*/
func main() {
	numbers := []int{2, 1, 7, 5, 9, 6, 3, 8, 4}
	selectsort(numbers,len(numbers))
	for _, n := range numbers {
		println(n)
	}
}

func selectsort(numbers []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0 ; i < n ; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		numbers[i],numbers[min] = numbers[min], numbers[i]
	}
}

	