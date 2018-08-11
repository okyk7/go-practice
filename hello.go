package main

import (
	"fmt" // 前にアルファベットつけるとalias、_つけると使わないやつ、つけると直接参照
)

var a string = "a"
var b, c, d string = "b", "c", "d"
var (
	e string = "e"
	f        = "f"
)

// mainパッケージのmain関数から実行される
func main() {
	fmt.Println("hello world")
	fmt.Println(b)
	fmt.Println(f)

	// :=するとvar宣言いらないし、内容から型推論される
	// 関数内限定らしい
	message2 := "hello world"
	fmt.Println(message2)

	const Hello string = "hello"
	fmt.Println(Hello)

	var i int // 初期化されないやつは決まってる値で初期化される
	fmt.Println(i)

	j, k := 1, 2
	if j < k {
		fmt.Println("kがでかい")
	} else {
		fmt.Println("jがでかい")
	}

	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	var x int
	// whileとかない
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// 無限ループ
	x = 0
	for {
		if x > 10 {
			break
		}
		x++
	}
}
