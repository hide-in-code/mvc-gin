package tool

import (
	"fmt"
)

//便捷打印
func Dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
