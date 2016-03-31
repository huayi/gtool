package list

import (
    "reflect"
    "fmt"
)

type List struct {
    Array []interface{}
}

func (this *List) Add(o ...interface{}) {
    this.Array = append(this.Array, o...)
}

func (this *List) AddAll(cons interface{}) *List {
    val := reflect.ValueOf(cons)
    switch reflect.TypeOf(cons).Kind() {
    case reflect.Slice, reflect.Array:
        for i := 0; i < val.Len(); i++ {
            this.Add(val.Index(i).Interface())
        }
    }
    return this
}

func (this *List) Delete(o interface{}) bool {
    for i, elem := range this.Array {
        if equal(elem, o) {
            this.Array = append(this.Array[:i], this.Array[i+1:]...)
            return true
        }
    }
    return false
}

func (this *List) Replace(i int, o interface{}) interface{} {
    //todo 提取indexCheck函数用来检查下标越界,并抛出panic
    if i >=0 && i < this.Len() {
        origin := this.Array[i]
        this.Array[i] = o
        return origin
    } else {
        fmt.Println("err: index out of bounds i:", i)
        return nil
    }

}

func (this *List) RmDuplicate() {
    for i:=0; i<this.Len(); i++ {
        for j:=i+1; j<this.Len(); j++ {
            if equal(this.Array[i], this.Array[j]) {
                this.Array = append(this.Array[:j], this.Array[j+1:]...)
            }
        }
    }

}

func (this *List) Len() int {
    return len(this.Array)
}

func (this *List) IsEmpty() bool {
    return len(this.Array) == 0
}

func (this *List) Contain(o interface{}) bool {
    return this.IndexOf(o) >= 0

}

func (this *List) IndexOf(o interface{}) int {
    for i, elem := range this.Array {
        if equal(elem, o) {
            return i
        }
    }
    return -1
}

func (this *List) Clear() {
    for i:=0; i<len(this.Array); i++ {
        this.Array[i] = nil
    }
}

func equal(o1, o2 interface{}) bool {
    if e, ok := o1.(Equaler); ok {
        return e.Equal(o2)
    } else {
        return o1 == o2
    }
}