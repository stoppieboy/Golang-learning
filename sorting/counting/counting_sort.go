// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func main() {
    var size int
    fmt.Print("Enter the size of array: ")
    fmt.Scanf("%d\n", &size)
    ar := make([]int, size)
    for i:=0; i < size; i++ {
        fmt.Scanf("%d\n", &ar[i])
    }
    sort(ar)
}

func sort(ar []int) []int{
    var k int
    n:=len(ar)
    for _,x := range ar{
        if x > k{
            k = x
        }
    }
    count := make([]int, k+1)
    res := make([]int, n)
    for i:=0; i < n; i++{
        count[ar[i]]++
    }
    for i:=1; i < k+1; i++ {
        count[i] += count[i-1]
    }
    // var temp int
    for i:=n-1; i >= 0; i--{
        res[count[ar[i]]-1] = ar[i]
        count[ar[i]] -= 1
    }
    fmt.Println(res)
    return res
}
