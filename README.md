Short lived, in memory RRDs
---------------------------------

Stupid simple RRDs that are only kept in memory.

Example
-------

    package main

	import (
		"fmt"
		"github.com/daniellarence/go-memrrd"
	)

	func main() {
		example_rrd := RRD{}
		example_rrd.New("Example")
		example_rrd.Add(1)
		fmt.Println(example_rrd)
	}
