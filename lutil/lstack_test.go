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

import "testing"


func TestStack(t *testing.T) {

	s := NewStack()

	if !s.Empty() {
  	t.Errorf("Stack should be empty")
	}

	s.Push(1972)
	s.Push(1974)
	s.Push(1977)

	if s.Len() != 3 {
  	t.Errorf("Stack should contain 3 elements")
	}

	i := s.Pop()

	if i.(int) != 1977 {
  	t.Errorf("Pop should return 1977 instead of %d",i.(int))
	}

	if s.Len() != 2 {
  	t.Errorf("Stack should contain 2 elements")
	}

	i = s.Pop()

	if i.(int) != 1974 {
  	t.Errorf("Pop should return 1974 instead of %d",i.(int))
	}

	i = s.Top()

	if i.(int) != 1972 {
  	t.Errorf("Top should return 1972 instead of %d",i.(int))
	}

	i = s.Top()

	if i.(int) != 1972 {
  	t.Errorf("Top should return 1972 instead of %d",i.(int))
	}

	i = s.Pop()

	if i.(int) != 1972 {
  	t.Errorf("Pop should return 1972 instead of %d",i.(int))
	}

	if !s.Empty() {
  	t.Errorf("Stack should be empty")
	}
}
