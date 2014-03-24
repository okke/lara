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

import "github.com/okke/lara/lutil"

type ObjectID uint64
type TypeID uint8

const (
	TYPE_OBJECT    = 100
	TYPE_PRIMITIVE = 101
)

type object struct {
	id ObjectID

	deleted bool

	refCount uint

	properties map[string]Object

	value Primitive

	listeners lutil.Set
}

type Typeable interface {
	Type() TypeID
}

type Object interface {
	Typeable

	Id() ObjectID
	Delete()

  Refer()
	Release()
	RefCount() uint

	Value() Primitive
	Link(name string, to Object)
	Get(name string) Object

	AddListener(listener Listener)
}

//
// --- [object constructors] ---
//

func newObject(id ObjectID) *object {
	o := new(object)
	o.id = id

	o.properties = make(map[string]Object)

	o.listeners = lutil.NewSet()

	return o
}

func newInt(id ObjectID, i Integer) *object {
	o := newObject(id)
	o.value = NewIntPrimitive(i)
	return o
}

func newFloat(id ObjectID, f Float) *object {
	o := newObject(id)
	o.value = NewFloatPrimitive(f)
	return o
}

func newString(id ObjectID, s string) *object {
	o := newObject(id)
	o.value = NewStringPrimitive(s)
	return o
}

func newList(id ObjectID) *object {
	o := newObject(id)
	o.value = NewListPrimitive()
	return o
}

//
// --- [object implements Typable] ---
//

func (self *object) Type() TypeID {
	if self.value != nil {
		return TYPE_PRIMITIVE
	} else {
		return TYPE_OBJECT
	}
}

//
// --- [object implements Object] ---
//

func (self *object) Id() ObjectID {
	return self.id
}

func (self *object) Delete() {
	if self.deleted {
		return
	}


	informDelete(self, self.listeners)
	self.listeners.RemoveAll()

	for key, el := range self.properties {
		el.Release()
		self.properties[key] = nil
	}
	self.value = nil

	self.deleted = true
}

func (self *object) Refer() {
  self.refCount++
}

func (self *object) Release() {
  self.refCount--
}

func (self *object) RefCount() uint {
  return self.refCount
}

func (self *object) Value() Primitive {
	return self.value
}

func (self *object) Link(name string, to Object) {

  if self.properties[name] != nil {
		self.properties[name].Release()
	}

	self.properties[name] = to

	to.Refer()

	to.AddListener(NewDeleteListener(func(o Object) {
		delete(self.properties, name)
	}))
}

func (self *object) Get(name string) Object {
	return self.properties[name]
}

func (self *object) AddListener(listener Listener) {
	self.listeners.Add(listener)
}
