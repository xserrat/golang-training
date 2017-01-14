## Packages

Capitalization matters:
 * Capitalize variable/function means visible outside (public)
 * Lowercase variable/function means not visible outside (private)

To import our packages we need to use the `Fully qualified name` as follow:

```go
    import (
        "fmt"
        
        "github.com/xserrat/golang-training/some-package"
    )
    
    func APublicFunction()  {    
        some-package.privateFunction()
    }
    
    func aPrivateFunction()  {
        fmt.Println(some-package.privateVariable)
        fmt.Println(some-package.PublicVariable)
    }
```

The `main` package is the package we can execute using `go run` command.