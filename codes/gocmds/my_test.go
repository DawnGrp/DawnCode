package main

import "testing"

func TestHello(t *testing.T) {

	a := 10
	b := 3

	t.Logf("a^b is %d", a^b)
}
