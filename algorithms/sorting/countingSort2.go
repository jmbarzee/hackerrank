package main
import "fmt"

func main() {
    var num, temp int
    fmt.Scanf("%v", &num)
    arr := make([]int, 100)
    for i := 0; i < num; i++ {
        fmt.Scanf("%v", &temp)
        arr[temp]++
    }
    for i := 0; i < 100; i++ {
        for arr[i] > 0 {
            fmt.Printf("%v ", i)
            arr[i]--
        }
    }
}