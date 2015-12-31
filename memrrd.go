package main

import (
	"fmt"
	"encoding/json"
	"time"
)

type RRD struct {
	Name  string     `json:"name"`
	Value []float64  `json:"value"`
	Lock  bool       `json:"lock"`
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

func tickAdd(r *RRD) {
	t := time.Now()
	now_epoch_int64 := t.Unix()
	now_epoch := float64(now_epoch_int64)
	fmt.Printf("Trying to add %d\n", now_epoch)
	for {
		if ! r.Lock {
			r.Lock = true
			r.Add(now_epoch)
			r.Lock = false
			break
		}
		fmt.Println("Failed to add new epoch")
	}

}

func main() {
	r := RRD{
		Name:   "name",
		Value: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}}
	go tickAdd(&r)
	fmt.Println(r.Json())
}
