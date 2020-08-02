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
