package data_stu

import "testing"

func BenchmarkAppendStringUseAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendStringUseAdd()
	}
}

func BenchmarkAppendStringUseJoin(b *testing.B) {
	for i := 0; i < b.N; i++  {
		AppendStringUseJoin()
	}
}

func BenchmarkAppendStringUseBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendStringUseBuffer()
	}
}
