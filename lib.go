package lib

import (
	"fmt"
	"go/build"
	"runtime"
)

func Run() {
	fmt.Println("lib:")
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("Toolchain version: %s\n", build.Default.ReleaseTags[len(build.Default.ReleaseTags)-1])
}
