package ormStu

import (
	"fmt"
	"testing"
)

func TestUser_Insert(t *testing.T) {
	u := &User{Name: "HelloWorld", Age: 24}

	u.Insert(u)
}

func BenchmarkUser_Insert(b *testing.B) {
	u := &User{Name: "HelloWorld", Age: 24}

	u.Insert(u)
}

func TestUser_GetAll(t *testing.T) {
	if all, err := new(User).GetAll(); err != nil {
		panic(err)
	} else {
		fmt.Println(all)
	}
}

func BenchmarkUser_GetAll(b *testing.B) {
	if all, err := new(User).GetAll(); err != nil {
		panic(err)
	} else {
		fmt.Println(all)
	}
}
