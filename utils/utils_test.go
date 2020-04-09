package utils

import "testing"

func TestThisIsTest(t *testing.T) {
	test := ThisIsTest()
	if test != ""{
		t.Fatal("Something wrong...")
	}
	t.Log("This is log")
}
