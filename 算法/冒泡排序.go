package main
/**
* 比较相邻两数大小进行排序
* 外层循环控制排序趟数，内层循环比较相邻数字大小完成本趟排序
* 空间复杂度：O(1) 时间复杂度:最好O(n) 最坏O(n②) 平均O(n②)
*/
func main() {
	numbers := []int{2, 1, 7, 5, 9, 6, 3, 8, 4}
	sort(numbers)
	for _, n := range numbers {
		println(n)
	}

}

func sort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				//temp := numbers[j]
				//numbers[j] = numbers[j+1]
				//numbers[j+1] = temp
				numbers[j], numbers[j + 1] = numbers[j + 1], numbers[j]
			}
		}
	}
}
