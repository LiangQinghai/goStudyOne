package data_stu

import (
	"bytes"
	"strings"
	"unsafe"
)

func CopyTmp(by []byte) string{
	return *(*string)(unsafe.Pointer(&by))
}

func AppendStringUseAdd()  string{
	s := ""
	for i := 0; i < 100000 ; i ++ {
		s += string(i)
	}
	return s
}

func AppendStringUseJoin() string {
	s := make([]string, 100000)
	for i := 0; i < 100000 ; i ++ {
		s[i] = string(i)
	}
	return strings.Join(s, "")
}

func AppendStringUseBuffer()  string{
	var b  bytes.Buffer
	b.Grow(100000)
	for i := 0; i < 100000; i++ {
		b.WriteString(string(i))
	}
	return b.String()
}