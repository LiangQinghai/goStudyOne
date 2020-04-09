package main

type Map map[string]string

type iMap Map

func main() {

	mm := make(map[string]string, 10)
	mm["one"] = "one"

	var mmap Map = mm

	println(mmap)

	// 强制转换 (Type) (var)
	var mimap iMap = (iMap) (mmap)

	print(mimap)

}