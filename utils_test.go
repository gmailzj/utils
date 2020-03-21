package utils_test

import (
    "fmt"
    "testing"
    "utils"
)

func TestIntSumSum(t *testing.T)  {
    var a = 3
    var b =4
    res :=utils.IntSum(a,b)
    fmt.Printf("%d 与%d之和:为%d",a,b,res)
    if res != 8{
        t.Error("error")
    }
}