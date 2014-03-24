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

type space struct {
	baseId ObjectID
	count  uint64

	all map[ObjectID]Object
}

type Space interface {
	Create() Object
	CreateInt(i Integer) Object
	CreateFloat(f Float) Object
	CreateString(s string) Object
	CreateList() Object
	Count() uint64

	Find(id ObjectID) Object
}

func NewSpace() Space {
	space := new(space)
	space.all = make(map[ObjectID]Object)
	return space
}

//
// --- [space implements Space] ---
//

func (self *space) create(creator func() Object) Object {
	self.baseId++
	self.count++
	o := creator()

  self.all[o.Id()] = o

	o.AddListener(self)
	return o
}

func (self *space) Create() Object {
	return self.create(func() Object {
		return newObject(self.baseId)
	})
}

func (self *space) CreateInt(i Integer) Object {
	return self.create(func() Object {
		return newInt(self.baseId, i)
	})
}

func (self *space) CreateFloat(f Float) Object {
	return self.create(func() Object {
		return newFloat(self.baseId, f)
	})
}

func (self *space) CreateString(s string) Object {
	return self.create(func() Object {
		return newString(self.baseId, s)
	})
}

func (self *space) CreateList() Object {
	list := self.create(func() Object {
		return newList(self.baseId)
	})

	list.AddListener(NewDeleteListener(func(o Object) {
		list.Value().ListValue().Each(func(val Object) {
			val.Release()
		})
	}))

  return list
}

func (self *space) Count() uint64 {
	return self.count
}

func (self *space) Find(id ObjectID) Object {
	return self.all[id]
}

//
// --- [space implements ObjectListener] ---
//

func (self *space) OnDelete(object Object) {
	self.count--
	self.all[object.Id()] = nil
}
