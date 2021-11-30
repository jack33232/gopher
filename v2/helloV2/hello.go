package helloV2

import (
	"fmt"

	"github.com/jack33232/gopher/hello"
)

func Hello() {
	hello.Hello()
	fmt.Println("Extra Hello from Version 2 package")
}
