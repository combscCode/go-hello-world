// This package is Chris's first attempt at writing go code.
package main

import (
	"fmt"
	"github.com/combscCode/go-hello-world/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
