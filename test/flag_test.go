package test

import (
	"flag"
	"testing"
)

func test_flag(t *testing.T) {
	var v bool
	flag.BoolVar(&v, "v", false, "print the version")

}
