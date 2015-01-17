package main

import "fmt"

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

func main() {
	a := []int{7, 7, 11, 7 9, 7, 7, 6, 5, 7, 3, 7, 1}
	t := find(a)
    fmt.Println(find(a))
    
	resort(a)
	fmt.Println(a)
}
