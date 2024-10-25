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
	ss.Append(&S{
		A: 2,
	})
	ss.Append(&S{
		A: 3,
	})

	// 获取长度
	fmt.Println("Length:", ss.Len())

	// 获取元素
	if val, ok := ss.Get(1); ok {
		fmt.Println("Element at index 1:", val)
	}

	// 删除元素
	if ss.Remove(0) {
		fmt.Println("Removed element at index 0")
	}

	// 输出最终 Slice
	fmt.Println("Final Slice length:", ss.Len())

	ss.Range(func(index int, value *S) {
		fmt.Println(index, value)
	})
}
