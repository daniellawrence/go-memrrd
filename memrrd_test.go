package memrrd

import (
	"testing"
)

func Test_memrrd_set_name(t *testing.T) {
	rrd := RRD{
		Name: "test",
	}
	if rrd.Name != "test" {
		t.Error("Failed to set name")
	}
}

func Test_memrrd_default_is_empty(t *testing.T) {
	rrd := RRD{
		Name: "test",
		Value: []float64{1.0}}
	expected := "{\"name\":\"test\",\"value\":[1]}"

	if rrd.Json() != expected {
		t.Error("Matching problem")
		t.Error(expected)
		t.Error(rrd.Json())
	}
}

func Test_memrrd_add1(t *testing.T) {
	rrd := RRD{
		Name: "test",
		Value: make([]float64, 3)}
	expected := "{\"name\":\"test\",\"value\":[0,0,0]}"

	if rrd.Json() != expected {
		t.Error("Matching problem")
		t.Error(expected)
		t.Error(rrd.Json())
	}
}
