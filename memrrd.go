package memrrd

import (
	"fmt"
	"encoding/json"
)
type RRD struct {
	Name  string     `json:"name"`
	Value []float64  `json:"value"`
}

func (self *RRD) Add(new_value interface{}) {
	copy(self.Value[1:], self.Value[0:])
	if x, ok := new_value.(float64); ok {
		self.Value[0] = x
		return
	}
}

func (self *RRD) Json() string {
	b, err := json.Marshal(self)
	if err != nil {
		return "{\"error\": \"Failed to convert RRD to json\"}"
	}
	return string(b)
}

func main() {
	r := RRD{
		Name:   "1",
		Value: []float64{1.0}}
	fmt.Println(r.Json())
}
