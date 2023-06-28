# Format

```go
package main

import (
 "fmt"

 "github.com/bluebrown/format"
)

func main() {
 fmt.Println(format.ListOr("apply", "orange", "kiwi"))
 // apple, orange or kiwi
}
```

If you want to change the `or` term, set the global varibale `OrTerm`.

```go
func init() {
 format.OrTerm = "oder"
}
```
