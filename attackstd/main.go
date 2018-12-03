package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args)!=3 {
		panic("usage: attackstd url N")
	}

	url := os.Args[1]
	N,_ := strconv.Atoi(os.Args[2])

	request := url+"/?guess=awrongguess"

	timings := make([]int64,N)

	for i:=0;i<N;i++ {
		start := time.Now()

		http.Get(request)

		end := time.Now()

		if i%100==0 {
			fmt.Println("count",i)
		}

		timings[i]=end.Sub(start).Nanoseconds()
	}

	var total int64 = 0

	for _,t := range timings {
		total = total + t
	}

	avg := total/int64(N)

	total = 0

	for _,t := range timings {
		a := t - avg
		total = (a*a)
	}

	s := math.Sqrt(float64(total/int64(N-1)))

	fmt.Println("avg time ",avg,"nanos","std",s,"nanos")
}

