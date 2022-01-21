// main_test.go

package main

import (
	"io/ioutil"
	"testing"
)

func TestConfig(t *testing.T) {
	_, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		t.Errorf("ERROR: Test - Missing config file")
	}
}
