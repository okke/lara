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
package lvm

import "github.com/okke/lara/lom"

const (
  OPCODE_CREATE_OBJECT=1
  OPCODE_CREATE_INT=2
  OPCODE_CREATE_FLOAT=3
  OPCODE_CREATE_STRING=4
  OPCODE_CREATE_LIST=5
)

type instruction struct {
  opcode int8
	params []interface{}
}

type Instruction interface {
}

func newInstruction(code int8, params []interface{}) *instruction {
  i := new(instruction)
	i.opcode = code
	i.params = params

	return i;
}

func NewCreateObject() Instruction {
	return newInstruction(OPCODE_CREATE_OBJECT, nil)
}

func NewCreateInt(i lom.Integer) Instruction {
	return newInstruction(OPCODE_CREATE_INT, []interface{}{i})
}

func NewCreateFloat(f lom.Float) Instruction {
	return newInstruction(OPCODE_CREATE_FLOAT, []interface{}{f})
}

func NewCreateString(s string) Instruction {
	return newInstruction(OPCODE_CREATE_STRING, []interface{}{s})
}

func NewCreateList() Instruction {
	return newInstruction(OPCODE_CREATE_LIST, nil)
}

