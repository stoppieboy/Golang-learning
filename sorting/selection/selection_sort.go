// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func main() {
    var size int
    fmt.Print("Enter the size of array: ")
    fmt.Scanf("%d\n",&size)
    var ar = make([]int, size)
    fmt.Println("Enter the elements: ")
    for i:=0; i < size; i++ {
        fmt.Scanf("%d\n",&ar[i])
    } 
    // ar := []int{4, 5, 2, 8, 1}
    fmt.Println(ar)
    ar = sort(ar)
    fmt.Print(ar)
}

func sort(ar []int) []int{
    start := 0
    var min int
    n := len(ar)
    for start < n-1{
        min = start;
        for i:=start+1; i < len(ar); i++ {
            if ar[i] < ar[min] {
                min = i
            }
        }
        if ar[min] < ar[start] {
            fmt.Print(ar[min]," ")
            ar[min], ar[start] = ar[start], ar[min]
            fmt.Println(ar[min])
        }
        start++
    }
    fmt.Println(ar)
    return ar
}
