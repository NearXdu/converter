# converter

## :blue_book: Introduction

golang type conversions tool based on functional programming




## :hammer: Quick Start


```golang
package main

import (
	"fmt"
	c "github.com/NearXdu/converter"
)

func main() {
	t1 := "hello,world"
	r1 := c.Convert2StringSet(c.SplitWith(","))(t1)
	fmt.Printf("%v",r1)
}
```
