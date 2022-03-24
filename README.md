
<h1><p align="center">gopy</p></h1>

<p align="center">
  <a href="https://github.com/sonlamho/gopy/actions">
    <img src="https://img.shields.io/github/workflow/status/sonlamho/gopy/Go?style=flat-square" alt="Github Actions">
  </a>
  <img src="https://img.shields.io/github/go-mod/go-version/sonlamho/gopy?style=flat-square">
  <a href="https://github.com/sonlamho/gopy/releases">
    <img src="https://img.shields.io/github/release/sonlamho/gopy/all.svg?style=flat-square">
  </a>
  <a href="https://goreportcard.com/report/github.com/sonlamho/gopy">
    <img src="https://goreportcard.com/badge/github.com/sonlamho/gopy#">
  </a>
  <a href="https://pkg.go.dev/github.com/sonlamho/gopy"><img src="https://pkg.go.dev/badge/github.com/sonlamho/gopy.svg" alt="Go Reference"></a>
</p>

## Features

- Basic functional tools: `Map`, `Filter`, `Reduce`
- Reducing functions: `Sum`, `Min`, `Max`, `All`, `Any`
- Variadic version of some functions: `VarSum`, `VarMin`, `VarMax`, `VarAll`, `VarAny`

## Example Usage

Before using `gopy` in your project, get the package:
```go get github.com/sonlamho/gopy@latest```


Sample code below:
```golang
package main

import (
	"fmt"
	"github.com/sonlamho/gopy"
)

var Print = fmt.Println

func main() {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	add_1half := func(x int) float64 { return float64(x) + 0.5 }
	is_even := func(x int) bool { return x%2 == 0 }

	Print(seq)
	Print(gopy.Map(add_1half, seq))
	Print(gopy.Filter(is_even, seq))
	Print(gopy.Sum(seq))
	Print("---")
	Print(gopy.Min(seq))  // -> 1
	Print(gopy.VarMin(100, 50, 9, 42))  // -> 9
	Print(gopy.VarMin(0.5, 0.7))  // -> 0.5
}
```
Output:
```
[1 2 3 4 5 6 7 8 9 10]
[1.5 2.5 3.5 4.5 5.5 6.5 7.5 8.5 9.5 10.5]
[2 4 6 8 10]
55
---
1
9
0.5
```
