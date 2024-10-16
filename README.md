# TUtils

TUtils is a simple library with various helpers that can be used in terminal applications

## Installation

```bash
go get github.com/Perkovec/tutils
```

## Methods

### GetSize

Returns the size of the terminal window

```go
package main

import (
    "fmt"
    "github.com/Perkovec/tutils"
)

func main() {
    fmt.Println(tutils.GetSize()) // example: Size{Rows: 24, Columns: 80}
}
```