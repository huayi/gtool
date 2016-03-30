package list

type List struct {
    Array []interface{}
}

func (this *List) Add(o ...interface{}) {
    this.Array = append(this.Array, o...)
}

func (this *List) Delete(o interface{}) bool {
    if e, ok := o.(Equaler); ok {
        for i, elem := range this.Array {
            if ok && e.Equal(elem){
                this.Array = append(this.Array[:i], this.Array[i+1:]...)
                return true
            }
        }
    } else {
        for i, elem := range this.Array {
            if o == elem {
                this.Array = append(this.Array[:i], this.Array[i+1:]...)
                return true
            }
        }
    }

    return false
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
    if e, ok := o.(Equaler); ok {
        for i, elem := range this.Array {
            if e.Equal(elem) {
                return i
            }
        }
    } else {
        for i, elem := range this.Array {
            if o == elem {
                return i
            }
        }
    }
    return -1
}

func (this *List) Clear() {
    for i:=0; i<len(this.Array); i++ {
        this.Array[i] = nil
    }
}