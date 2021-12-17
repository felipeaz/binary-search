# binary-search
Binary search implementation in Go

Binary search is used to find an element into an ordered array. The algorithm works by reducing the analysis range in half

Algorithm complexity: O(log n) or O(1) if we find the target on the first iteration

### Benchmark Test
On the project root, run the following command: 

`` go test -count ${n} -bench=.`` 

replace `${n}` with the number of times you'd like to run the test

### Unit Test
On the project root, run the following command:

`` go test -v .`` 
