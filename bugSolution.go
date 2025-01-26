```Go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var m sync.Mutex // Changed to sync.Mutex
        var count int

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m.Lock()
                        count++
                        m.Unlock()
                        fmt.Printf("Goroutine %d: count = %d\n", i, count)
                }(i)
        }

        wg.Wait()
        fmt.Printf("Final count: %d\n", count)
}
```