# hands-on-high-performance-with-go
Hands-On High Performace with Go


In this section, you will learn why performance in computer science is importmant. You will alse learn why performance is important in the Go language. Moving on, you will learn about data structures and algorithms, concurrency, STL algorithm equivalents, and the matrix and vector computations in Go.

The chapters in this section include the following:

* Chapter 1, introduction to Performace in Go.

* Chapter 2, Data Structures and Algorithms.
  -- Benchmark example: 
  --- Run: go test -bench=. -benchtime 2s -count 2 -benchmem -cpu 4

  goos: darwin
  goarch: amd64
  pkg: github.com/mrojasb2000/hands-on-high-performance-with-go/chapter2/benchmark/benchmark_test
  BenchmarkHello-4        24789703                88.4 ns/op            32 B/op          1 allocs/op
  PASS
  ok      github.com/mrojasb2000/hands-on-high-performance-with-go/chapter2/benchmark/benchmark_test   4.712s

  GOOS     : Operating System
  GOARCH   : Architecture
  -4       : The number process of GOMAXPROCS that were used to execute.
  24789703 : The number of times out loop ran to gather the necesary data.
  88.4 ns/op : The speed per loop during out test.
  PASS     : Indicates the end state of our benchmark run.
  

* Chapter 3, Understanding Concurrency.

* Chapter 4, STL Algorithm Equivalents in Go.

* Chapter 5, Matrix and Vector Computation in Go.
