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

import "testing"

func TestCreateObject(t *testing.T) {

	space := NewSpace()

	object := space.Create()

	if object.Id() != 1 {
		t.Errorf("Object id of first object is not 1")
	}

	if object.Type() != TYPE_OBJECT {
		t.Errorf("Object has not valid type");
	}

	object2 := space.Create()

	if object2.Id() != 2 {
		t.Errorf("Object id of second object is not 2")
	}

	count := space.Count()

	if count != 2 {
		t.Errorf("There should be 2 objects")
	}

}

func TestCreatePrimitive(t *testing.T) {
	space := NewSpace()

	i := space.CreateInt(74839)

	if i.Type() != TYPE_PRIMITIVE {
		t.Errorf("Integer Object has not valid type");
	}

	p := i.Value()
	if p == nil {
		t.Errorf("Integer Object does not have a primitive value")
	}

	if p.Type() != TYPE_INTEGER {
		t.Errorf("Integer Object has not valid primitive type");
	}

	if p.IntValue() != 74839 {
		t.Errorf("Integer Object has wrong value. Expected %d, got %d", 74839, p.IntValue())
	}

	f := space.CreateFloat(74839.4738)

	if f.Type() != TYPE_PRIMITIVE {
		t.Errorf("Float Object has not valid type");
	}

	p = f.Value()
	if p == nil {
		t.Errorf("Float Object does not have a primitive value")
	}

	if p.Type() != TYPE_FLOAT {
		t.Errorf("Integer Object has not valid primitive type");
	}

	if p.FloatValue() != 74839.4738 {
		t.Errorf("Float Object has wrong value. Expected %f, got %f", 74839.4738, p.FloatValue())
	}

	s := space.CreateString("soup")

	if s.Type() != TYPE_PRIMITIVE {
		t.Errorf("String Object has not valid type");
	}

	p = s.Value()
	if p == nil {
		t.Errorf("String Object does not have a primitive value")
	}

	if p.Type() != TYPE_STRING {
		t.Errorf("String Object has not valid primitive type");
	}

	if p.StringValue() != "soup" {
		t.Errorf("String Object has wrong value. Expected soup, got %s", p.StringValue())
	}

	count := space.Count()

	if count != 3 {
		t.Errorf("There should be 3 objects")
	}

}

func TestCreateList(t *testing.T) {
	space := NewSpace()

	l := space.CreateList()

	if l.Type() != TYPE_PRIMITIVE {
		t.Errorf("List Object has not valid type");
	}

	p := l.Value()
	if p == nil {
		t.Errorf("List Object does not have a primitive value")
	}

	if p.Type() != TYPE_LIST {
		t.Errorf("List Object has not valid primitive type")
	}

  arr := p.ListValue()

	if (arr.Size() != 0) {
		t.Errorf("List Object should be empty")
	}

  first := space.Create()
	arr.Add(first)

	if (arr.Size() != 1) {
		t.Errorf("List Object should contain 1 object instead of %d", arr.Size())
	}

	arr.Add(space.CreateInt(913203))

	if (arr.Size() != 2) {
		t.Errorf("List Object should contain 2 objects instead of %d", arr.Size())
	}

	object := space.Create()

	arr.Add(object)

	if (arr.Size() != 3) {
		t.Errorf("List Object should contain 3 objects instead of %d", arr.Size())
	}

	t.Logf("List:%v",arr)

	object.Delete()

	if (arr.Size() != 2) {
		t.Errorf("List Object should contain 2 objects instead of %d", arr.Size())
	}

	t.Logf("List:%v",arr)

	object = space.Create()
	arr.Add(object)
	arr.Add(space.CreateInt(123))

	t.Logf("List:%v",arr)

	object.Delete()

	t.Logf("List:%v",arr)

	if (arr.Size() != 3) {
		t.Errorf("List Object should contain 3 objects instead of %d", arr.Size())
	}

	first.Delete()

	t.Logf("List:%v",arr)

	if (arr.Size() != 2) {
		t.Errorf("List Object should contain 2 objects instead of %d", arr.Size())
	}

}

func TestDeleteObject(t *testing.T) {

	space := NewSpace()

	object := space.Create()

	if space.Count() != 1 {
		t.Errorf("There should be 1 object")
	}

	object.Delete()

	if space.Count() != 0 {
		t.Errorf("There should be no objects")
	}

	object.Delete()

	if space.Count() != 0 {
		t.Errorf("There should be no objects, found %d", space.Count())
	}
}

func TestLinkObject(t *testing.T) {

	space := NewSpace()

	object1 := space.Create()
	object2 := space.Create()
	id2 := object2.Id()


	object1.Link("likes", object2)
	linkedId := object1.Get("likes").Id()

	if linkedId != id2 {
		t.Errorf("Link should refer to object id %d instead of %d",id2,linkedId)
	}

	object2.Delete()
	linkedObjectAfterDelete := object1.Get("likes")
	if linkedObjectAfterDelete != nil {
		t.Errorf("Link should be removed after deletion of linked object")
	}


}

func TestRefCounting(t *testing.T) {

	space := NewSpace()

	object1 := space.Create()
	object2 := space.Create()

	object1.Link("likes", object2)

	if object2.RefCount() != 1 {
		t.Errorf("Creating a link should increase the reference count")
	}

	object1.Link("likes", space.Create())

	if object2.RefCount() != 0 {
		t.Errorf("Unlinking should decrease the reference count")
	}

	object3 := space.Create()
	object1.Link("likes", object3)
	object1.Delete()

	if object3.RefCount() != 0 {
		t.Errorf("Deleting an object should decrease the reference count")
	}

	object4 := space.Create()
	l := space.CreateList()
	a := l.Value().ListValue()

	a.Add(object4)
	if object4.RefCount() != 1 {
		t.Errorf("Adding to a list should increase the reference count")
	}

	l.Delete()
	if object4.RefCount() != 0 {
		t.Errorf("Deleting a list should decrease the reference count")
	}

}

func TestFindById(t *testing.T) {
	space := NewSpace()

	id1 := space.Create().Id()
	id2 := space.Create().Id()

	object1 := space.Find(id1)
	object2 := space.Find(id2)

	if object1 == nil {
		t.Errorf("object 1 not found")
	}

	if object2 == nil {
		t.Errorf("object 2 not found")
	}

	if object1 == object2 {
		t.Errorf("objects can not be the same")
	}

	object1.Delete()
	object1 = space.Find(id1)

	if object1 != nil {
		t.Errorf("object 1 should be deleted")
	}
}
