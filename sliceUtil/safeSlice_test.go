package sliceUtil

import (
	"fmt"
	"testing"
)

type S struct {
	A int
}

func TestSafeSlice(t *testing.T) {
	ss := &SafeSlice[S]{}

	// 添加元素
	ss.Append(&S{
		A: 1,
	})
	ss.Insert(0, &S{
		A: 2,
	})
	ss.Append(&S{
		A: 3,
	})
	ss.Set(1, &S{
		A: 4,
	})
	// 删除元素
	if ss.RemoveAt(2) {
		fmt.Println("Removed element at index 0")
	}

	// 输出最终 Slice
	fmt.Println("Final Slice length:", ss.Len())

	for i, v := range ss.Values() {
		fmt.Println(i, v)
	}

	ss.Clear()

	a := ss.Values()[0:]
	fmt.Println(a)
}
