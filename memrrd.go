package memrrd

import (
	"fmt"
)

type RRD struct {
	name  string
	value [10]int
}

func (self *RRD) New(name string)  {
	self.name = name
}

func (self *RRD) prepend(new_value int) {
	copy(self.value[1:], self.value[0:])
	self.value[0] = new_value
}

// func main() {
//	cpu_rrd := RRD{}
//	cpu_rrd.New("cpu")
//	cpu_rrd.prepend(1)
//	fmt.Println(cpu_rrd)
// }
