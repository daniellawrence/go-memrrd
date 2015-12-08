package memrrd

import (
	"testing"
)

func Test_memrrd_set_name(t *testing.T) {
	rrd := RRD{}
	rrd.New("Example")
	// example_rrd.Add(1)
	if rrd.name != "Example" {
		t.Error("Failed to set name")
	}
}

func Test_memrrd_default_is_empty(t *testing.T) {
	rrd := RRD{}
	rrd.New("Example")
	if rrd.value != [10]int{0,0,0,0,0,0,0,0,0,0} {
		t.Error("Failed to init value")
	}
}

func Test_memrrd_add1(t *testing.T) {
	rrd := RRD{}
	rrd.New("Example")
	rrd.Add(1)
	if rrd.value != [10]int{1,0,0,0,0,0,0,0,0,0} {
		t.Error("Failed to add value")
	}
	rrd.Add(2)
	if rrd.value != [10]int{2,1,0,0,0,0,0,0,0,0} {
		t.Error("Failed to add value")
	}
}
