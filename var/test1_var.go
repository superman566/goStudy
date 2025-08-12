package main

import "fmt"

// 局部变量 → 优先 :=

// 包级变量 → 用 var

// 常量 → 用 const

// 需要强调类型的 → 用 var（即使能推断，也显式写类型）

// go
// Copy
// Edit
// var buf bytes.Buffer  // 明确是 bytes.Buffer 类型
// 避免滥用零值 → 如果需要默认值以外的初始化，就用 :=

// 📌 社区共识总结

// 简短、临时 → :=

// 全局、零值初始化、需要类型信息 → var

// 永不变 → const

func main() {
	// method 1
	var a int
	fmt.Println("a=", a)

	// method 2
	var b int = 100
	fmt.Println("b=", b)

	// method 3
	var c = 100
	fmt.Println("c=", c)
	fmt.Printf("type of c = %T \n", c)

	// method 4 := short variable declaration
	e := 200
	fmt.Println("e=", e)

	var x, y = 300, 400
	fmt.Println("x=", x, "y=", y)
	var ( 
		vv int = 100
		jj= true
	)
}
