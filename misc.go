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

func main() {
    fmt.Println(Kf(257))
}
