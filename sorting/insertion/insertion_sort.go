package main

import "fmt"

func main() {
    var size int
    fmt.Print("Enter the size of array: ")
    fmt.Scanf("%d\n",&size)
    ar := make([]int, size)
    for i:=0; i<size; i++ {
        fmt.Scanf("%d\n",&ar[i])
    }
    // ar := []int{5,4,6,2,7}
    fmt.Println(ar)
    ar = sort(ar)
    fmt.Println(ar)
}

func sort(ar []int) []int{
    fmt.Println("here")
    n := len(ar)
    var j int
    var temp int
    for i:=0; i<n; i++ {
        j = i+1
        temp = i
        for j < n && temp >= 0 && ar[j] < ar[temp] {
            ar[j], ar[temp] = ar[temp], ar[j]
            j--
            temp--
        }
    }
    return ar
}