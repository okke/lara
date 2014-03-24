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

/*
 forked from github.com/bemasher/stack.go
*/

package lutil

type stack struct {
	top *elem
	size int
}

type elem struct {
	value interface{}
	next *elem
}

type Stack interface {
	Empty() bool
	Len() int
	Push(value interface{})
	Pop() interface{}
	Top() interface{}
}

func NewStack() Stack {
	return new(stack)
}

func (s *stack) Empty() bool {
	return s.size == 0
}

func (s *stack) Len() int {
	return s.size
}

func (s *stack) Push(value interface{}) {
	s.top = &elem{value, s.top}
	s.size++
}

func (s *stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *stack) Top() interface{} {
	if s.size > 0 {
		return s.top.value
	}
	return nil
}
									 
