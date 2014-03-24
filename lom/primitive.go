/*
Copyright (c) 2014, Okke van 't Verlaat

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package lom

const (
	TYPE_UNKNOWN = 0
	TYPE_INTEGER = 1
	TYPE_FLOAT   = 2
	TYPE_STRING  = 3
	TYPE_LIST    = 4
)

type Integer int64
type Float float64

type primitive struct {
	primType TypeID

	value interface{}
}

type List interface {
  Size() int
  Add(val Object) 

	Each(f func(val Object))
}


type Primitive interface {
	Typeable

	IntValue() Integer
	FloatValue() Float
	StringValue() string
	ListValue() List
}

//
// --- [primitive constructors] ---
//

func NewIntPrimitive(val Integer) Primitive {
	obj := new(primitive)
	obj.primType = TYPE_INTEGER
	obj.value = val
	return obj
}

func NewFloatPrimitive(val Float) Primitive {
	obj := new(primitive)
	obj.primType = TYPE_FLOAT
	obj.value = val
	return obj
}

func NewStringPrimitive(s string) Primitive {
	obj := new(primitive)
	obj.primType = TYPE_STRING
	obj.value = s
	return obj
}

func NewListPrimitive() Primitive {
	obj := new(primitive)
	obj.primType = TYPE_LIST
	obj.value = make([]Object,0,16)
	return obj
}

//
// --- [primitive implements Typable] ---
//

func (self *primitive) Type() TypeID {
	return self.primType
}

//
// --- [primitive implements Primitive] ---
//

func (self *primitive) IntValue() Integer {
  return self.value.(Integer)
}

func (self *primitive) FloatValue() Float {
  return self.value.(Float)
}

func (self *primitive) StringValue() string {
  return self.value.(string)
}

func (self *primitive) ListValue() List {
  return self
}

//
// --- [primitive implements List] ---
//

func (self *primitive) Size() int {
  return len(self.value.([]Object))
}

func (self *primitive) Add(val Object) {

  self.value = append(self.value.([]Object), val)

	val.Refer()

	i := len(self.value.([]Object))

	val.AddListener(NewDeleteListener(func(o Object) {
	  a := self.value.([]Object)
		copy(a[i-1:], a[i:])
		a[len(a)-1] = nil
		self.value = a[:len(a)-1]
	}))
}

func (self *primitive) Each(f func(val Object)) {
	for _,el := range self.value.([]Object) {
  	f(el)
	}
}
