package main

import (
	"fmt"
	"encoding/json"
)
type RRD struct {
	Name  string     `json:"name"`
	Value []float64  `json:"value"`
}

func (self *RRD) Add(new_value float64) {
	copy(self.Value[1:], self.Value[0:])
	self.Value[0] = new_value
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
		Name:   "name",
		Value: []float64{1.0, 2.0}}
	r.Add(3.0)
	fmt.Println(r.Json())
}
