package main

import "fmt"

func kf(start, end, d int) (r int) {
    mid := (start + end) >> 1

    switch {
    case start == mid || mid * mid == d:
        r = mid
    case mid * mid < d:
        r = kf(mid, end, d)
    case mid * mid > d:
        r = kf(start, mid, d)
    }
    return r
}

func Kf(d int) (i int) {
    return kf(0, d, d)
}

func shaizi(d int) []int {
    r := make([]int, d)
    for i := 2; i <= d; i++ {
        r[i-1] = i
    }

    var rd []int
    for i := 2; i * i <= d; i++ {
        for j := i + i; j <= d; j += i {
            r[j - 1] = 0
        }
    }

    for _, v := range r {
        if v != 0 {
            rd = append(rd, v)
        }
    }

    return rd
}

func main() {
    //fmt.Println(Kf(257))
    fmt.Println(shaizi(125))
}
