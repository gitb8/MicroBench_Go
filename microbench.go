package main

// add packages for bench:

import (
	"fmt"
	"runtime"
	"time"
	// add imports for bench:
)

// add variables for bench:

func main() {
	// set run repetitions:
	const RUNS int = 65536

	var start time.Time = time.Now()
	var first time.Duration = time.Since(start)
	var second time.Duration = time.Since(start)

	// add local variables for bench:

	var a int64 = 0

	start = time.Now()
	for i := 0; i < RUNS; i++ {
		// 1st run

		a++
		// a += 1		// reverse test

		// bench function
		func1()
		// func2()		// reverse test
	}
	first = time.Since(start)

	// reset bench variables:

	a = 0

	// run garbage collector
	runtime.GC()

	start = time.Now()
	for i := 0; i < RUNS; i++ {
		// 2nd run

		a += 1
		// a++			// reverse test

		// bench function
		func2()
		// func1()		// reverse test
	}
	second = time.Since(start)

	// print result nanoseconds
	// fmt.Println("1st run: ", float32(first)/float32(RUNS), " ns")
	// fmt.Println("2nd run: ", float32(second)/float32(RUNS), " ns")

	// print result microseconds
	fmt.Println("1st run: ", float32(first)/float32(RUNS)/1000.0, " µs")
	fmt.Println("2nd run: ", float32(second)/float32(RUNS)/1000.0, " µs")

	// print result milliseconds
	// fmt.Println("1st run: ", float32(first)/float32(RUNS)/1.0e6, " ms")
	// fmt.Println("2nd run: ", float32(second)/float32(RUNS)/1.0e6, " ms")

	// print result seconds
	// fmt.Println("1st run: ", float32(first)/float32(RUNS)/1.0e9, " s")
	// fmt.Println("2nd run: ", float32(second)/float32(RUNS)/1.0e9, " s")
}

// add functions for bench:

func func1() {} // demo function

func func2() {} // demo function
