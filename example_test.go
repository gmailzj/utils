package utils_test

import (
    "fmt"
    "utils"
)

/**
如果为函数添加Example, 直接在example_test.go文件中添加函数func Example{要提供示例的函数名}
如果为方法添加Example, 直接在example_test.go文件中添加函数func Example{方法接受者结构体类型名}_{要提供示例的方法名}
如果为类型添加Example,直接在example_test.go文件中添加函数fun Example{类型名}
 */

func ExampleGetEmptyString(){

    str := utils.GetEmptyString()
    fmt.Println(str)
}

func ExampleMaxSlice() {
    d := []int{1,2,3,4}
    max := utils.MaxSlice(d)
    fmt.Println(max)
    //Output:
    //4
}