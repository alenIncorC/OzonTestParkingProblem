# Parking Problem

---

###command to run program 
- go run main.go

###command to run benchmarks:
#### benchmarks with memory quick overview
- go test -bench . -benchmem
  
  ````
    goos: darwin 
    goarch: arm64 
    BenchmarkParkingProblem-8       510670261                2.078 ns/op           0 B/op          0 allocs/op
  ````
  
---
#### pprof
- go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1
- go tool pprof parkingProblem.test cpu
```
(pprof) list recParkOuter
Total: 980ms
ROUTINE ======================== _/Users/alenamanbekov/Documents/projects/beeline-projects/parkingProblem.recParkOuter in /Users/alenamanbekov/Documents/projects/beeline-projects/parkingProblem/main.go
     390ms      390ms (flat, cum) 39.80% of Total
         .          .     38:
         .          .     39:   i++
         .          .     40:   return recParkInner(s, acc, i)
         .          .     41:}
         .          .     42:
      80ms       80ms     43:func recParkOuter(p []pair, b int, inc int, r func(s []pair, acc int, i int) int) int {
     130ms      130ms     44:   if len(p) == inc {
     180ms      180ms     45:           return b
         .          .     46:   }
         .          .     47:
         .          .     48:   res := r(p[inc:], 0, 0)
         .          .     49:   if b < res {
         .          .     50:           b = res

```

- go tool pprof parkingProblem.test mem.out
```
(pprof) list recParkOuter
Total: 2.41MB
```
