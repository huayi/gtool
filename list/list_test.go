package list

import (
    "testing"
    "fmt"
)

func TestDelete(t *testing.T) {
    list := new(List)
    list.Delete("1")
    fmt.Println(list)
}

func TestIndexOf(t *testing.T) {
    list := new(List)
    list.Add("2","3","4","3")
    i := list.IndexOf("3")
    fmt.Println(i)
}