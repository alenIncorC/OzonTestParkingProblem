package main

import (
	"fmt"
	"math/rand"
	"time"
)

type pair [2]int64

func (p pair) car() int64 {
	return p[0]
}

func (p pair) cdr() int64 {
	return p[1]
}

var seq []pair

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now())
	for i := 0; i < 100; i++ {
		seq = append(seq, pair{rand.Int63n(86400) + time.Now().UnixNano(), rand.Int63n(86400*7) + time.Now().UnixNano() })
	}
	fmt.Println(recParkOuter(seq, 0, 0, recParkInner))
}

func recParkInner(s []pair, acc int, i int) int {
	if len(s) == i {
		return acc
	}

	if s[0].cdr() >= s[i].car() && s[0].car() <= s[i].cdr() {
		acc++
	}

	i++
	return recParkInner(s, acc, i)
}

func recParkOuter(p []pair, b int, inc int, r func(s []pair, acc int, i int) int) int {
	if len(p) == inc {
		return b
	}

	res := r(p[inc:], 0, 0)
	if b < res {
		b = res
	}
	inc++
	return recParkOuter(p, b, inc, r)
}
