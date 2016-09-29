package main
import "fmt"

// the original task is here https://www.hackerrank.com/challenges/circular-array-rotation

func main() {
    var N, K, Q int
    // read the input
    fmt.Scanf("%v %v %v", &N, &K, &Q)
    ar := make([]int, N, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&ar[i])
    }
    
    if K > N {
        K = K - N * int(K/N)
    }
    
    ar = append(ar[N-K:],ar[:N-K]...)
    
    for i := 0; i < Q; i++{
        var j int
        fmt.Scan(&j)
        fmt.Println(ar[j])
    }
}
