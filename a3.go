
package main

import (
    "fmt"
    "sync"
)

func main() {
    var pool = sync.Pool{
        New: func() interface{} {
            return "new object"
        },
    }

    obj := pool.Get().(string)
    fmt.Println(obj)  // Output: new object

    pool.Put("reused object")
    obj = pool.Get().(string)
    fmt.Println(obj)  // Output: reused object
}