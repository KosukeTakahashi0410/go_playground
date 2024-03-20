package main

import "fmt"

// Hello world

func main() {
	// fmt.Println("Hello world")
	// fmt.Println(time.Now())

	// var 変数名　型　=　値

	var i int = 100
	fmt.Println(i)

	var s string = "Hello world"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t)
	fmt.Println(f)

	var (
		i2 int    = 200
		s2 string = "test"
	)

	fmt.Println(i2)
	fmt.Println(s2)

	// 暗黙的な定義
	// 変数名 := 値
	testvalue := "string type"
	fmt.Println(testvalue)
}
