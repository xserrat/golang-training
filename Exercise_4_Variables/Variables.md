## Variables

Two ways:

* Shorthand: (only used inside functions!!)

```go
package main

import "fmt"

func main() {
        name := "Xevi"
        fmt.Println(name)
}
```

* Variable with zero value assignment: (inside or outside functions)

```go
package main

import "fmt"

func main() {
        var name string
        
        name = "Xevi"
        fmt.Println(name)
}
```