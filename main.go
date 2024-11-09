package main

import (
	"fmt"
	"go_type_inference/parser"
)

func main() {
	input := `if 2 < 3 then (if true then x else y) else false`

	res := parser.Parse(input)
	fmt.Println(res)
}
