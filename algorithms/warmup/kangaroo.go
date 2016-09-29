package main
import "fmt"

func main() {
    var k1, s1, k2, s2 int
    fmt.Scanf("%v %v %v %v", &k1, &s1, &k2, &s2)
    
    fmt.Println(solve(k1,k2,s1,s2))
}

func solve(k1,k2,s1,s2 int) string {
    for k1 != k2 {
        k1 += s1
        k2 += s2
        if k1 > k2 && s1 >= s2 {
            return "NO"
        }
        if k2 > k1 && s2 >= s1{
            return "NO"
        }
    }
    return "YES"
}
