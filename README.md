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
