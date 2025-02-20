```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var m sync.Mutex
        var x int

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m.Lock()
                        temp := x
                        temp += i
                        x = temp
                        m.Unlock()
                }(i)
        }

        wg.Wait()
        fmt.Println("x:", x)
}
```