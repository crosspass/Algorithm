package main

import "fmt"

// 程序之美-寻找水王
func find(d []int) int {
	var i, t int
	for _, v := range d {
		if i == 0 {
			t = v
		}
		if t == v {
			i++
		} else {
			i--
		}
	}
	return t
}

// 归并排序
func resort_d(d []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) >> 1
	resort_d(d, start, mid)
	mid++
	resort_d(d, mid, end)
	dl := make([]int, (mid - start))
	copy(dl, d[start:mid])
	var k int
	for k = 0; k < len(dl) && mid <= end; start++ {
		if dl[k] <= d[mid] {
			d[start] = dl[k]
			k++
		} else {
			d[start] = d[mid]
			mid++
		}
	}

	for k < len(dl) {
		d[start] = dl[k]
		start++
		k++
	}
}

func resort(d []int) {
	resort_d(d, 0, len(d)-1)
}

func find_one(d int) (count int) {
	for d != 0 {
		if t := d % 10; t == 1 {
			count++
		}
		d /= 10
	}
	return
}

// 程序之美-寻找1的个数
func find_ones(d int) (sum int) {
	for i := 0; i <= d; i++ {
		sum += find_one(i)
	}
	return
}

func main() {
	a := []int{7, 7, 11, 7, 9, 7, 7, 6, 5, 7, 3, 7, 1}
	fmt.Println(find(a))

	resort(a)
	fmt.Println(a)
	fmt.Println(find_ones(12))
}
