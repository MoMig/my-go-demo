package main

/**
* 每一趟排序选择一个基准元素，一趟排序结束，形成以基准元素分界的两部分，左边数比右边小，然后以相同方式递归左右两边
* 空间复杂度：O(nlog②n) 时间复杂度:最好O(nlog②n) 最坏O(n2) 平均O(nlog②n)
*/
func main() {
	numbers := []int{2, 1, 7, 5, 9, 6, 3, 8, 4}
	quicksort(numbers, 0, len(numbers) - 1)
	for _, n := range numbers {
		println(n)
	}

}

func quicksort(numbers []int, left int, right int) {
	if left >= right {
		return
	}

	temp := numbers[left]
	i := left
	j := right
	for i < j {
		//先找右边小于基准的数
		for i < j && numbers[j] >= temp {
			j--
		}
		//再找左边大于基准的数
		for i < j && numbers[i] <= temp {
			i++
		}
		//左右存在需要交换位置的元素
		if i < j {
			temp := numbers[i]
			numbers[i] = numbers[j]
			numbers[j] = temp
		}
	}
	numbers[left] = numbers[i]
	//基准数移到中间
	numbers[i] = temp
	quicksort(numbers, left, i-1)
	quicksort(numbers, i+1, right)

}
