package main

import (
	"fmt"

	_base94 "github.com/BEN00262/base94/base94"
)

func main() {
	base94 := _base94.New()

	if val, err := base94.Encode(789); err == nil {
		fmt.Println(val)

		fmt.Println(base94.Decode(val))
	}
}
