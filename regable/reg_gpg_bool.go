///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Wed Oct 05 2016 23:12:04]
// Generate from:
//   [github.com/vipally/gx/regable/reg.gp]
//   [github.com/vipally/gx/regable/reg.gpg] [bool]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct  5 2016 22:08:02]
// Version : 2.9.0
// 
///////////////////////////////////////////////////////////////////

package regable

import (
	"bytes"
	"fmt"
	"sync"
	"github.com/vipally/gx/consts"
	"github.com/vipally/gx/errors"
	xmath "github.com/vipally/gx/math"
)

const (
	default_bool_reg_cnt = 1024
)

var (
	g_bool_rgr_id_gen, _         = amath.NewRangeUint32(g_invalid_id+1, g_invalid_id, g_max_reger_id)
	errid_bool_id, _  = aerr.Reg("BoolId error")
	errid_bool_obj, _ = aerr.Reg("Bool object error")
)

var (
	g_bool_reger_list []*BoolReger
)

func init() {
	reg_show(ShowAllBoolRegers)
}

//new reger
func NewBoolReger(name string) (r *BoolReger, err error) {
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	if err = check_lock(); err != nil {
		return
	}
	id := g_invalid_id
	if id, err = g_bool_rgr_id_gen.Inc(); err != nil {
		return
	}
	p := new(BoolReger)
	if err = p.init(name); err == nil {
		p.reger_id = uint8(id)
		r = p
		g_bool_reger_list = append(g_bool_reger_list, p)
	}
	return
}

func MustNewBoolReger(name string) (r *BoolReger) {
	if reg, err := NewBoolReger(name); err != nil {
		panic(err)
	} else {
		r = reg
	}
	return
}

//show all regers
func ShowAllBoolRegers() string {
	var buf bytes.Buffer
	s := fmt.Sprintf("[BoolRegers] count:%d", len(g_bool_reger_list))
	buf.WriteString(s)
	for _, v := range g_bool_reger_list {
		buf.WriteString(consts.NewLine)
		buf.WriteString(v.String())
	}
	return buf.String()
}

//reger object
type BoolReger struct {
	reger_id uint8
	name     string
	id_gen   amath.RangeUint32
	reg_list []*_boolRecord
}

func (me *BoolReger) init(name string) (err error) {
	me.name = name
	if err = me.id_gen.Init(g_invalid_id+1, g_invalid_id,
		g_invalid_id+default_bool_reg_cnt); err != nil {
		return
	}
	me.reg_list = make([]*_boolRecord, 0, 0)
	return
}

//set max reg count at a reger
func (me *BoolReger) MaxReg(max_regs uint32) (rmax uint32, err error) {
	if err = verify_max_regs(max_regs); err != nil {
		return
	}
	cur, min, _ := me.id_gen.Get()
	if err = me.id_gen.Init(cur, min, g_invalid_id+max_regs); err != nil {
		return
	}
	rmax = me.id_gen.Max()
	return
}

//reg a value
func (me *BoolReger) Reg(name string, val bool) (r BoolId, err error) {
	r = BoolId(g_invalid_id)
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	id := g_invalid_id
	if err = check_lock(); err != nil {
		return
	}
	if id, err = me.id_gen.Inc(); err != nil {
		return
	}
	p := me.new_rec(name, val)
	p.id = id
	me.reg_list = append(me.reg_list, p)
	r = BoolId(MakeRegedId(uint32(me.reger_id), id))
	return
}

func (me *BoolReger) MustReg(name string, val bool) (r BoolId) {
	if reg, err := me.Reg(name, val); err != nil {
		panic(err)
	} else {
		r = reg
	}
	return
}

//show string
func (me *BoolReger) String() string {
	var buf bytes.Buffer
	s := fmt.Sprintf("[BoolReger#%d: %s] ids:%s",
		me.reger_id, me.name, me.id_gen.String())
	buf.WriteString(s)
	for i, v := range me.reg_list {
		v.lock.RLock()
		s = fmt.Sprintf("\n#%d [%s]: %v",
			uint32(i)+g_invalid_id+1, v.name,
			v.val)
		v.lock.RUnlock()
		buf.WriteString(s)
	}
	return buf.String()
}

type _boolRecord struct {
	name string
	val  bool
	id   uint32
	lock sync.RWMutex
}

func (me *BoolReger) new_rec(name string, val bool) (r *_boolRecord) {
	r = new(_boolRecord)
	r.name = name
	r.val = val
	return
}

type BoolId regedId

func (cp BoolId) get() (rg *BoolReger, r *_boolRecord, err error) {
	idrgr, id := regedId(cp).ids()
	idregidx, idx := idrgr-g_invalid_id-1, id-g_invalid_id-1
	if idrgr == g_invalid_id || !g_bool_rgr_id_gen.InCurrentRange(idrgr) {
		err = aerr.New(errid_bool_id)
	}
	rg = g_bool_reger_list[idregidx]
	if id == g_invalid_id || !rg.id_gen.InCurrentRange(id) {
		err = aerr.New(errid_bool_id)
	}
	r = rg.reg_list[idx]
	return
}

//check if valid
func (cp BoolId) Valid() (rvalid bool) {
	if _, _, e := cp.get(); e == nil {
		rvalid = true
	}
	return
}

//get value
func (cp BoolId) Get() (r bool, err error) {
	_, rc, e := cp.get()
	if e != nil {
		return r, e
	}
	return rc.Get()
}

//get value with out error, if has error will cause panic
func (cp BoolId) GetNoErr() (r bool) {
	_, rc, e := cp.get()
	if e != nil {
		panic(e.Error())
	}
	return rc.GetNoErr()
}

//set value
func (cp BoolId) Set(val bool) (r bool, err error) {
	_, rc, e := cp.get()
	if e != nil {
		return r, e
	}
	return rc.Set(val)
}

//reverse bool value(as a switch)
func (cp BoolId) Reverse() (r bool, err error) {
	_, rc, e := cp.get()
	if e != nil {
		return r, e
	}
	return rc.Reverse()
}

//get reger_id and real_id
func (cp BoolId) Ids() (reger_id, real_id uint32) {
	return regedId(cp).ids()
}

//show string
func (cp BoolId) String() (r string) {
	idrgr, id := regedId(cp).ids()
	_, rc, err := cp.get()
	if err != nil {
		r = fmt.Sprintf("invalid BoolId#(%d|%d)", idrgr, id)
	} else {
		r = rc.String()
	}
	return
}

//get name

func (cp BoolId) Name() (r string, err error) {
	_, rc, e := cp.get()
	if e == nil {
		r, err = rc.Name()
	} else {
		err = e
	}
	return
}


//get as object for fast access
func (cp BoolId) Oject() (r BoolObj) {
	_, rc, e := cp.get()
	if e == nil {
		r.obj = rc
	}
	return
}

//get name

func (me *_boolRecord) Name() (r string, err error) {
	if me != nil {
		me.lock.RLock()
		defer me.lock.RUnlock()
		r = me.name
	} else {
		err = aerr.New(errid_bool_obj)
	}
	return
}


//get value
func (me *_boolRecord) Get() (r bool, err error) {
	if me != nil {
		me.lock.RLock()
		defer me.lock.RUnlock()
		r = me.val
	} else {
		err = aerr.New(errid_bool_obj)
	}
	return
}

//get value without error,if has error will cause panic
func (me *_boolRecord) GetNoErr() (r bool) {
	r0, err := me.Get()
	if err != nil {
		panic(err.Error())
	}
	r = r0
	return
}

//set value
func (me *_boolRecord) Set(val bool) (r bool, err error) {
	if nil != me {
		me.lock.Lock()
		defer me.lock.Unlock()
		me.val = val
		r = val
	} else {
		err = aerr.New(errid_bool_obj)
	}
	return
}

//reverse on bool value
func (me *_boolRecord) Reverse() (r bool, err error) {
	if nil != me {
		me.lock.Lock()
		defer me.lock.Unlock()
		me.val = !me.val
		r = me.val
	} else {
		err = aerr.New(errid_bool_obj)
	}
	return
}

//get as Id
func (me *_boolRecord) Id() (r BoolId) {
	if me != nil {
		r = BoolId(me.id)
	}
	return
}

//show string
func (me *_boolRecord) String() (r string) {
	if me != nil {
		idrgr, id := regedId(me.id).ids()
		me.lock.RLock()
		defer me.lock.RUnlock()
		r = fmt.Sprintf("Bool#(%d|%d|%s)%v", idrgr, id, me.name, me.val)
	} else {
		r = fmt.Sprintf("invalid bool object")
	}
	return
}

//object of reged value,it is more efficient to access than Id object
type BoolObj struct {
	obj *_boolRecord
}

//check if valid
func (cp BoolObj) Valid() (rvalid bool) {
	return cp.obj != nil
}

//get value
func (cp BoolObj) Get() (r bool, err error) {
	return cp.obj.Get()
}

//get value against error,if has error will cause panic
func (cp BoolObj) GetNoErr() (r bool) {
	return cp.obj.GetNoErr()
}

//set value
func (cp BoolObj) Set(val bool) (r bool, err error) {
	return cp.obj.Set(val)
}

//reverse bool object
func (cp BoolObj) Reverse() (r bool, err error) {
	return cp.obj.Reverse()
}

//show string
func (cp BoolObj) String() (r string) {
	return cp.obj.String()
}

//get name

func (cp BoolObj) Name() (r string, err error) {
	return cp.obj.Name()
}


//get as Id
func (cp BoolObj) Id() (r BoolId) {
	return cp.obj.Id()
}

//reg and return an object agent
func (me *BoolReger) RegO(name string, val bool) (r BoolObj, err error) {
	id, e := me.Reg(name, val)
	if e == nil {
		r = id.Oject()
	} else {
		err = e
	}
	return
}

func (me *BoolReger) MustRegO(name string, val bool) (r BoolObj) {
	if reg, err := me.RegO(name, val); err != nil {
		panic(err)
	} else {
		r = reg
	}
	return
}
