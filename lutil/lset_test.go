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

func TestSetAdd(t *testing.T) {

	set := NewSet()

	set.Add(1)
	set.Add(3)

  if !set.Contains(1) {
  	t.Errorf("Set should contain value 1");
  }

  if set.Contains(2) {
  	t.Errorf("Set should not contain value 2");
  }

  if !set.Contains(3) {
  	t.Errorf("Set should contain value 3");
  }
}

func TestSetRemove(t *testing.T) {

	set := NewSet()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Remove(1)
	set.Remove(3)

  if set.Contains(1) {
  	t.Errorf("Set should not contain value 1");
  }

  if !set.Contains(2) {
  	t.Errorf("Set should contain value 2");
  }

  if set.Contains(3) {
  	t.Errorf("Set should not contain value 3");
  }
}

func TestSetRemoveAll(t *testing.T) {

	set := NewSet()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.RemoveAll()

  if set.Contains(1) {
  	t.Errorf("Set should not contain value 1");
  }

  if set.Contains(2) {
  	t.Errorf("Set should not contain value 2");
  }

  if set.Contains(3) {
  	t.Errorf("Set should not contain value 3");
  }
}

func TestSetEach(t *testing.T) {

	set := NewSet()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	count := 0
	set.Each(func(val interface{}) {
	  count++
	})

  if count != 3 {
  	t.Errorf("Set should contain 3 values, found %d",count);
  }

	set.Remove(2)

	count = 0
	set.Each(func(val interface{}) {
	  count++
	})

  if count != 2 {
  	t.Errorf("Set should contain 2 values, found %d",count);
  }

	set.Add(1)
	set.Add(2)
	set.Add(3)

	count = 0
	set.Each(func(val interface{}) {
	  count++
	})

  if count != 3 {
  	t.Errorf("Set should contain 3 values, found %d",count);
  }

}
