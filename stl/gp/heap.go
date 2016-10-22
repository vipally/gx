package gp

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
//
//
///*   <----This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
//
// This is exactly not a real go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

//import...

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

import (
	dumy_fmt "fmt"
)

type GOGPHeapElem int

func (this GOGPHeapElem) Less(o GOGPHeapElem) bool {
	return this < o
}

func (me GOGPHeapElem) Show() string {
	return dumy_fmt.Sprintf("%d", me)
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//container object
type GOGPHeapNamePrefixHeap struct {
	b      []GOGPHeapElem //data buffer
	limitN int            //if limitN>0, heap size must<=limitN
	maxTop bool           //if top is max value
}

//push heap value
func (this *GOGPHeapNamePrefixHeap) PushHeap(b []GOGPHeapElem, v GOGPHeapElem) []GOGPHeapElem {
	b = append(b, v)
	this.adjustUp(b, len(b)-1, v)
	return b
}

//pop heap top
func (this *GOGPHeapNamePrefixHeap) PopHeap(b []GOGPHeapElem) (h []GOGPHeapElem, top GOGPHeapElem, ok bool) {
	l := len(b)
	if ok = l > 0; ok {
		top = b[0]
		if l > 1 {
			b[0], b[l-1] = b[l-1], b[0]
			this.adjustDown(b[:l-1], 0, b[0])
		}
		h = b[:l-1]
	}
	return
}

//check if b is a valid heap
func (this *GOGPHeapNamePrefixHeap) CheckHeap(b []GOGPHeapElem) bool {
	for i := len(b) - 1; i > 0; i-- {
		p := this.parent(i)
		if !this.cmpV(b[i], b[p]) {
			return false
		}
	}
	return true
}

//adjust heap to select a proper hole to set v
func (this *GOGPHeapNamePrefixHeap) adjustDown(b []GOGPHeapElem, hole int, v GOGPHeapElem) {
	size := len(b)
	//#if Gogp_ImproveHeap
	//try to improve STL's adjust down algorithm
	//adjust heap to select a proper hole to set v
	for l := this.lchild(hole); l < size; l = this.lchild(hole) {
		c := l                                              //index that need compare with hole
		if r := l + 1; r < size && !this.cmpV(b[r], b[l]) { //let the most proper child to compare with v
			c = r
		}
		if this.cmpV(b[c], v) { //v is the most proper root, finish adjust
			break
		} else { //c is the most proper root, swap with hole, and continue adjust
			b[hole], hole = b[c], c
		}
	}
	b[hole] = v //put v to last hole
	//#else
	//C++ stl's adjust down algorithm
	//it seems to cost more move, to get probably less cmpare
	for l := this.lchild(hole); l < size; l = this.lchild(hole) {
		c := l                                              //index that need to be new root
		if r := l + 1; r < size && !this.cmpV(b[r], b[l]) { //let the most proper child to compare with v
			c = r
		}
		b[hole], hole = b[c], c
	}
	this.adjustUp(b, hole, v) //adjust up from leaf hole
	//#endif
}

//adjust heap to select a proper hole to set v
func (this *GOGPHeapNamePrefixHeap) adjustUp(b []GOGPHeapElem, hole int, v GOGPHeapElem) {
	for hole > 0 {
		if parent := this.parent(hole); !this.cmpV(v, b[parent]) {
			b[hole], hole = b[parent], parent
		} else {
			break
		}
	}
	b[hole] = v //put v to last hole
}

//make b as a heap
func (this *GOGPHeapNamePrefixHeap) MakeHeap(b []GOGPHeapElem) {
	if l := len(b); l > 1 {
		for i := l / 2; i >= 0; i-- {
			this.adjustDown(b, i, b[i])
		}
	}
}

//reverse order of b
func (this *GOGPHeapNamePrefixHeap) Reverse(b []GOGPHeapElem) {
	l := len(b) - 1
	for i := l / 2; i >= 0; i-- {
		b[i], b[l-i] = b[l-i], b[i]
	}
}

//sort slice use heap algorithm
func (this *GOGPHeapNamePrefixHeap) SortHeap(b []GOGPHeapElem) []GOGPHeapElem {
	this.MakeHeap(b)
	for t := b; len(t) > 1; {
		t, _, _ = this.PopHeap(t)
	}
	return b
}

//new object
func NewGOGPHeapNamePrefixHeap(capacity int, maxTop bool) (r *GOGPHeapNamePrefixHeap) {
	r = &GOGPHeapNamePrefixHeap{}
	r.Init(capacity, 0, maxTop)
	return
}

//initialize
func (this *GOGPHeapNamePrefixHeap) Init(capacity, limitN int, maxTop bool) {
	if cap(this.b) < capacity {
		this.b = make([]GOGPHeapElem, 0, capacity)
	}
	this.b = this.b[:0]

	this.limitN = limitN
	this.maxTop = maxTop
}

//heap slice
func (this *GOGPHeapNamePrefixHeap) Buffer() []GOGPHeapElem {
	return this.b
}

//push
func (this *GOGPHeapNamePrefixHeap) Push(v GOGPHeapElem) {
	if this.limitN > 0 && this.Size() >= this.limitN {
		if top, ok := this.Top(); ok && this.cmpV(top, v) {
			this.Pop()
		}
	}
	this.b = this.PushHeap(this.b, v)
}

//pop
func (this *GOGPHeapNamePrefixHeap) Pop() (top GOGPHeapElem, ok bool) {
	this.b, top, ok = this.PopHeap(this.b)
	return
}

//heap top
func (this *GOGPHeapNamePrefixHeap) Top() (top GOGPHeapElem, ok bool) {
	if ok = !this.Empty(); ok {
		top = this.b[0]
	}
	return
}

//cmpare value
func (this *GOGPHeapNamePrefixHeap) cmpV(c, p GOGPHeapElem) (ok bool) {
	//#if GOGP_HasLess==true
	if this.maxTop {
		ok = !p.Less(c)
	} else {
		ok = !c.Less(p)
	}
	//#else
	if this.maxTop {
		ok = !(p < c)
	} else {
		ok = !(c < p)
	}
	//#endif
	return
}

//get parent index
func (this *GOGPHeapNamePrefixHeap) parent(idx int) int {
	return (idx - 1) / 2
}

//get left child index
func (this *GOGPHeapNamePrefixHeap) lchild(idx int) int {
	return 2*idx + 1
}

//get right child index
//func (this *GOGPHeapNamePrefixHeap) rchild(idx int) int {
//	return 2*idx + 2
//}

//func (this *GOGPHeapNamePrefixHeap) children(idx int) (l, r int) {
//	l = 2*idx + 1
//	r = l + 1
//	return
//}

//size
func (this *GOGPHeapNamePrefixHeap) Size() int {
	return len(this.b)
}

//empty
func (this *GOGPHeapNamePrefixHeap) Empty() bool {
	return this.Size() == 0
}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
