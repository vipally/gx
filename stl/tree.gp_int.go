///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Wed Nov 23 2016 22:36:23]
// Generate from:
//   [github.com/vipally/gx/stl/gp/tree.gp]
//   [github.com/vipally/gx/stl/stl.gpg] [tree_int]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct 24 2016 20:25:45]
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////

//this file defines a template tree structure just like file system

package stl

////////////////////////////////////////////////////////////////////////////////

//tree strture
type IntTree struct {
	root *IntTreeNode
}

//new container
func NewIntTree() *IntTree {
	p := &IntTree{}
	return p
}

//tree node
type IntTreeNode struct {
	int
	children IntTreeNodeSortSlice
}

func (this *IntTreeNode) Less(right *IntTreeNode) (ok bool) {

	ok = this.int < right.int

	return
}

func (this *IntTreeNode) SortChildren() {
	this.children.Sort()
}

func (this *IntTreeNode) Children() []*IntTreeNode {
	return this.children.Buffer()
}

//add a child
func (this *IntTreeNode) AddChild(v int, idx int) (child *IntTreeNode) {
	n := &IntTreeNode{int: v}
	return this.AddChildNode(n, idx)
}

//add a child node
func (this *IntTreeNode) AddChildNode(node *IntTreeNode, idx int) (child *IntTreeNode) {
	this.children.Insert(node, idx)
	return node
}

//cound of children
func (this *IntTreeNode) NumChildren() int {
	return this.children.Len()
}

//get child
func (this *IntTreeNode) GetChild(idx int) (child *IntTreeNode, ok bool) {
	child, ok = this.children.Get(idx)
	return
}

//remove child
func (this *IntTreeNode) RemoveChild(idx int) (child *IntTreeNode, ok bool) {
	child, ok = this.children.Remove(idx)
	return
}

//create a visitor
func (this *IntTreeNode) Visitor() (v *IntTreeNodeVisitor) {
	v = &IntTreeNodeVisitor{}
	v.push(this, -1)
	return
}

//get all node data
func (this *IntTreeNode) All() (list []int) {
	list = append(list, this.int)
	for _, v := range this.children.Buffer() {
		list = append(list, v.All()...)
	}
	return
}

//tree node visitor
type IntTreeNodeVisitor struct {
	node         *IntTreeNode
	parents      []*IntTreeNode
	brotherIdxes []int
	//visit order: this->child->brother
}

func (this *IntTreeNodeVisitor) push(n *IntTreeNode, bIdx int) {
	this.parents = append(this.parents, n)
	this.brotherIdxes = append(this.brotherIdxes, bIdx)
}

func (this *IntTreeNodeVisitor) pop() (n *IntTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n, bIdx = this.tail()
		this.parents = this.parents[:l-1]
		this.brotherIdxes = this.brotherIdxes[:l-1]
	}
	return
}

func (this *IntTreeNodeVisitor) tail() (n *IntTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n = this.parents[l-1]
		bIdx = this.brotherIdxes[l-1]
	}
	return
}

func (this *IntTreeNodeVisitor) depth() int {
	return len(this.parents)
}

func (this *IntTreeNodeVisitor) update_tail(bIdx int) bool {
	l := len(this.parents)
	if l > 0 {
		this.brotherIdxes[l-1] = bIdx
		return true
	}
	return false
}

func (this *IntTreeNodeVisitor) top_right(n *IntTreeNode) (p *IntTreeNode) {
	if n != nil {
		l := n.children.Len()
		for l > 0 {
			this.push(n, l-1)
			n = n.children.MustGet(l - 1)
			l = n.children.Len()
		}
		p = n
	}
	return
}

//visit next node
func (this *IntTreeNodeVisitor) Next() (ok bool) {
	if this.node != nil { //check if has any children
		if this.node.children.Len() > 0 {
			this.push(this.node, 0)
			this.node = this.node.children.MustGet(0)
		} else {
			this.node = nil
		}
	}
	for this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, bIdx := this.tail()
		if bIdx < 0 { //ref parent
			this.node = p
			this.pop()
		} else if bIdx < p.children.Len()-1 { //next brother
			bIdx++
			this.node = p.children.MustGet(bIdx)
			this.update_tail(bIdx)
		} else { //no more brothers
			this.pop()
		}
	}
	if ok = this.node != nil; ok {
		//do nothing
	}
	return
}

//visit previous node
func (this *IntTreeNodeVisitor) Prev() (ok bool) {
	if this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, _ := this.pop()
		this.node = this.top_right(p)
		if ok = this.node != nil; ok {
			//do nothing
		}
		return
	}

	if this.node != nil { //check if has any children
		p, bIdx := this.tail()
		if bIdx > 0 {
			bIdx--
			this.update_tail(bIdx)
			this.node = this.top_right(p.children.MustGet(bIdx))
		} else {
			this.node = p
			this.pop()
		}
	}
	if ok = this.node != nil; ok {
		//do nothing
	}
	return
}

//get node data
func (this *IntTreeNodeVisitor) Get() (data *int) {
	if nil != this.node {
		data = &this.node.int
	}
	return
}