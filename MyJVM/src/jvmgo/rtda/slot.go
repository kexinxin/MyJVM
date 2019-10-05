package rtda

import "jvmgo/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
