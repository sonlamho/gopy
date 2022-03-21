# gopy
makes Go feels more like Python


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
	Print(seq)
	Print(gopy.Map(func(x int) float64 { return float64(x) + 0.5 }, seq))
	Print(gopy.Filter(func(x int) bool { return (x%2 == 0) }, seq))
	Print(gopy.Sum(seq))

}
```
