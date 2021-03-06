# hands-on-high-performance-with-go
Hands-On High Performace with Go


In this section, you will learn why performance in computer science is importmant. You will alse learn why performance is important in the Go language. Moving on, you will learn about data structures and algorithms, concurrency, STL algorithm equivalents, and the matrix and vector computations in Go.

The chapters in this section include the following:

* Chapter 1, introduction to Performace in Go.

* Chapter 2, Data Structures and Algorithms.

  -- Benchmark example: <br>
  --- Run: go test -bench=. -benchtime 2s -count 2 -benchmem -cpu 4<br>
  --- Run: go test -bench=. -count 5 -cpu 1,2,4 > single.txt<br>

  ![Benchmark execution](images/Benchmark_run.png)

  GOOS     : Operating System<br>
  GOARCH   : Architecture<br>
  -4       : The number process of GOMAXPROCS that were used to execute.<br>
  24789703 : The number of times out loop ran to gather the necesary data.<br>
  88.4 ns/op : The speed per loop during out test.<br>
  PASS     : Indicates the end state of our benchmark run.<br>

  - Advantages
    - Surfaces potential problems before they become unwieldy
    - Helps developers have a deeper understanding of their code
    - Can identify potential bottlenecks in the design and data structures and algorithms stages

  - Drawbacks
    - Needs to be completed on a given cadence for meaningful results
    - Data collation can be difficult
    - Does not always yield a meaningful result for the problem at hand

 -- Benchmark comparation: 
 --- Run: benchstat -html -sort -delta single/single.txt multi/multi.txt > out.html

![Benchmark comparation](images/Benchstat_comparation.png)





* Chapter 3, Understanding Concurrency.

* Chapter 4, STL Algorithm Equivalents in Go.

* Chapter 5, Matrix and Vector Computation in Go.

