package main
/**
* 将待排序数组分为局部有序子数组和无序子数组，每次排序从无序子数组取第一个元素，从后向前与有序子数组进行比较，按照大小插入到合适位置
* 时间复杂度：O(n②) 最好O(n) 空间复杂度O(1)
*/
func main() {
	numbers := []int{2, 1, 7, 5, 9, 6, 3, 8, 4}
	insertsort(numbers,len(numbers))
	for _, n := range numbers {
		println(n)
	}
}

func insertsort(numbers []int, n int) {
	if n <= 1 {
		return
	}

	for i := 1 ; i < n ; i++ {
		v := numbers[i]
		j := i - 1
		for ; j >= 0 ;j-- {
			if numbers[j] > v {
				numbers[j + 1] = numbers[j]
			} else {
				break
			}
		}
		numbers[j + 1] = v
	}
}