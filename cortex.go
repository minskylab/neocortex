package neocortex

import "fmt"

var y = 2019

func hi() {

	hello := fmt.Sprintf("Hello from %d", y)
	fmt.Println(hello)
}
