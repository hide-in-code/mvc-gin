package tool

import (
	"encoding/json"
	"fmt"
)

//便捷打印
func Dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

//对象序列化
func Marshal(model interface{}) string {
	jsonStr, err := json.Marshal(model)
	if err != nil {
		fmt.Println("Marshal failed...")
	}

	return string(jsonStr)
}

//对象反序列化
func UnMarshal(model interface{}) string {
	jsonStr, err := json.Marshal(model)
	if err != nil {
		fmt.Println("Marshal failed...")
	}

	return string(jsonStr)
}
