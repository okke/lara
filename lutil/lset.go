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
package lutil


type set struct {
  values map[interface{}]bool
}

type Set interface {
  Add(val interface{})
	Remove(val interface{})
	RemoveAll()
	Contains(val interface{}) bool
	Each(f func(interface{}))
}

func NewSet() Set {
  s := new(set)
	s.values = make(map[interface{}]bool)
	return s
}

func (self *set) Add(val interface{}) {
	self.values[val] = true
}

func (self *set) Remove(val interface{}) {
	delete(self.values,val)
}

func (self *set) RemoveAll() {
	self.values = make(map[interface{}]bool)
}

func (self *set) Contains(val interface{}) bool {
  return self.values[val]
}

func (self *set) Each(f func(interface{})) {
  for val, _ := range self.values {
    f(val)
	}
}
