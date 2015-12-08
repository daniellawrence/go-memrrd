package memrrd

type RRD struct {
	name  string
	value [10]int
}

func (self *RRD) New(name string)  {
	self.name = name
}

func (self *RRD) Add(new_value int) {
	copy(self.value[1:], self.value[0:])
	self.value[0] = new_value
}
