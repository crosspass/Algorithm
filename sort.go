package main

import (
	"fmt"
	"time"
)

// 编程之美-寻找水王
func Find(d []int) int {
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

// 递归排序
func resortD(d []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) >> 1
	resortD(d, start, mid)
	mid++
	resortD(d, mid, end)
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

func Resort(d []int) {
	resortD(d, 0, len(d)-1)
}

func findOne(d int) (count int) {
	for d != 0 {
		if t := d % 10; t == 1 {
			count++
		}
		d /= 10
	}
	return
}

// 寻找数字1-编程之美
func FindOnes(d int) (sum int) {
	for i := 0; i <= d; i++ {
		sum += findOne(i)
	}
	return
}

// 寻找数字1v2版本-编程之美
func V2FindOnes(d int) (count int) {
	for i, tail := 1, 0; d != 0; {
		b := d % 10
		d /= 10
		switch {
		case b == 0:
			count += d * i
		case b == 1:
			count += d*i + tail + 1
		case b > 1:
			count += (d + 1) * i
		}
		tail += b * i
		i *= 10
	}
	return
}

// 精确表达浮点数-编程之美
func DisplayFloatNumber(f float32) string {
	return "todo"
}

func main() {
	a := []int{7, 7, 11, 7, 9, 7, 7, 6, 5, 7, 3, 7, 1}
	fmt.Println(Find(a))

	Resort(a)
	fmt.Println(a)
	t := time.Now()
	fmt.Println(FindOnes(33))
	fmt.Println(time.Now().Sub(t).Nanoseconds())
	t = time.Now()
	fmt.Println(V2FindOnes(33))
	fmt.Println(time.Now().Sub(t).Nanoseconds())
}
