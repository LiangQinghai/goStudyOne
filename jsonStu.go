package main

import (
	"encoding/json"
	"fmt"
)

type MyPerson struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {

	mp := MyPerson{Name:"Hello",Age:23}

	mpBytes01, e := json.Marshal(&mp)
	if e == nil {
		fmt.Println(string(mpBytes01))
	}else {
		panic(e)
	}

	var mpOne MyPerson
	jsonStr := "{\"name\":\"HelloWorld\",\"age\":24}"
	e = json.Unmarshal([]byte(jsonStr), &mpOne)
	if e == nil{
		fmt.Println(mpOne)
	}else {
		panic(e)
	}

}