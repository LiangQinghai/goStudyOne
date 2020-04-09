package main

import "fmt"

func main() {

	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["one"] = "one"

	fmt.Println(map1)

	map2 := map[string]string{
		"two": "two",
	}

	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Println(key, value)
	}

	//map是无序的，每次输出的结果顺序都可能不一样
	//要使map排序，首先将key放到一个切片中，然后斯通对该切片进行排序，最后再按照排序后的切片顺序取map中的值

}
