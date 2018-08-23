package main

import "testing"

func TestTransformName(t *testing.T) {
	s := "ops-test-4033419859-b6w75"

	name, err := transformName(s)
	if err != nil {
		t.Error("transform err", err)
	}
	if name != "ops-test" {
		t.Error("transform failed")
	}
}
