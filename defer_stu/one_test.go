package defer_stu

import (
	"fmt"
	"testing"
)

func BenchmarkNormalCall(b *testing.B) {
	for i := 0;i < b.N ;i++  {
		NormalCall()
	}
}

func BenchmarkDeferCall(b *testing.B) {
	a := [2]int{2, 3}
	c := [3]int{2, 3, 5}

	ints := append(a[0:], c[0:]...)

	println(ints)

	fmt.Println(a)
	for i := 0;i < b.N ;i++  {
		DeferCall()
	}
}
