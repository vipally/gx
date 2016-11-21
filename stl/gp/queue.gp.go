package gp

//#GOGP_FILE_BEGIN

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)

////////////////////////////////////////////////////////////////////////////////

//queue object
type GOGPGlobalNamePrefixQueue struct {
	//real data is [head,tail)
	//buffer d is cycle, that is to say, next(len(d)-1)=0, prev(0)=len(d)-1
	//so if tail<head, data is [head, end, 0, tail)
	head int
	tail int
	d    []GOGPValueType
}

//new object
func NewGOGPQueueNamePrefixQueue(bufSize int) *GOGPGlobalNamePrefixQueue {
	r := &GOGPGlobalNamePrefixQueue{}
	r.Init(bufSize)
	return r
}

//init
func (this *GOGPGlobalNamePrefixQueue) Init(bufSize int) {
	if nil == this.d {
		if bufSize <= 0 {
			bufSize = 8 //default buffer size
		}
		this.newBuf(bufSize)
	}
	this.Clear()
	return
}

//create new buffer
func (this *GOGPGlobalNamePrefixQueue) newBuf(bufSize int) {
	if bufSize > 0 {
		this.d = make([]GOGPValueType, bufSize, bufSize) //the same cap and len
	}
}

//clear
func (this *GOGPGlobalNamePrefixQueue) Clear() {
	this.head, this.tail = 0, 0
}

//push to back of queue
func (this *GOGPGlobalNamePrefixQueue) Push(v GOGPValueType) (ok bool) {
	if ok = true; ok {
		if nil == this.d { //init if needed
			this.Init(-1)
		}
		this.d[this.tail] = v
		if this.tail++; this.tail >= this.Cap() {
			this.tail = 0
			if this.tail == this.head { //tail catch up head, buffer full
				oldCap := this.Cap()
				d := this.d
				this.newBuf(oldCap * 2)
				h := copy(this.d, d[this.head:])
				t := copy(this.d[:h], d[:this.tail])
				this.head, this.tail = 0, h+t
			}
		}
	}
	return
}

//pop front of queue
func (this *GOGPGlobalNamePrefixQueue) Pop() (front GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		front = this.d[this.head]
		if this.head++; this.head >= this.Cap() && this.head != this.tail {
			this.head = 0
		}
	}
	return
}

//front data
func (this *GOGPGlobalNamePrefixQueue) Front() (front GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		front = this.d[this.head]
	}
	return
}

//back data
func (this *GOGPGlobalNamePrefixQueue) Back() (back GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		t := this.tail - 1
		if t < 0 {
			t = this.Cap() - 1
		}
		back = this.d[t]
	}
	return
}

//shrink data buffer if necessary
func (this *GOGPGlobalNamePrefixQueue) Shrink() (ok bool) {
	oldCap := this.Cap()
	oldSize := this.Size()
	if ok := oldCap > 8 && oldCap >= 3*oldSize; ok { //leave at least 8 elem space
		d := this.d
		this.newBuf(oldSize / 2)
		if this.tail >= this.head {
			copy(this.d, d[this.head:this.tail])
			this.tail -= this.head
			this.head = 0
		} else {
			h := copy(this.d, d[this.head:])
			t := copy(this.d[:h], d[:this.tail])
			this.head, this.tail = 0, h+t
		}
	}
	return
}

//func (this *GOGPDequeNamePrefixDeque) Sort() {}

//data buffer size
func (this *GOGPGlobalNamePrefixQueue) Cap() int {
	return len(this.d)
}

//size of queue
func (this *GOGPGlobalNamePrefixQueue) Size() (size int) {
	if this.tail >= this.head {
		size = this.tail - this.head
	} else {
		size = this.Cap() - this.head + this.tail
	}
	return
}

//if queue is empty
func (this *GOGPGlobalNamePrefixQueue) Empty() bool {
	return this.Size() == 0
}

////show
//func (this *GOGPQueueNamePrefixQueue) Show() string {
//	var b show_bytes.Buffer
//	b.WriteByte('[')
//	for i := this.head; i != this.tail; i++ {
//		if i >= this.Cap() {
//			i = 0
//		}
//		v = this.v[i]
//		b.WriteString(v.Show())
//		b.WriteByte(',')
//	}
//	if this.Depth() > 0 {
//		b.Truncate(b.Len() - 1) //remove last ','
//	}
//	b.WriteByte(']')
//	return b.String()
//}

//#GOGP_FILE_END
