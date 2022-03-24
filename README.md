
<h1><p align="center">gopy</p></h1>

<p align="center">
  <a href="https://github.com/sonlamho/gopy/actions">
    <img src="https://img.shields.io/github/workflow/status/sonlamho/gopy/Go?style=flat-square" alt="Github Actions">
  </a>
  <img src="https://img.shields.io/github/go-mod/go-version/sonlamho/gopy?style=flat-square">
  <a href="https://github.com/sonlamho/gopy/releases">
    <img src="https://img.shields.io/github/release/sonlamho/gopy/all.svg?style=flat-square">
  </a>
</p>


### Example Usage

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

}
```
Output:
```
[1 2 3 4 5 6 7 8 9 10]
[1.5 2.5 3.5 4.5 5.5 6.5 7.5 8.5 9.5 10.5]
[2 4 6 8 10]
55
```
