package array

import (
	"errors"
	"fmt"
)

type Array struct {
	values   []int
	length uint
}

func NewArray(leng uint) *Array {
	if leng == 0 {
		return nil
	}
	return &Array{
		values:   make([]int, leng, leng),
		length: 0,
	}
}

// give index it will return the element
func (a *Array) Find(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("Index out of range!")
	}
	return a.values[index], nil
}

func (a *Array) isIndexOutOfRange(index uint) bool {
	return index >= a.length
}

func (a *Array) Len() uint {
	return a.length
}

// insert values on index
func (a *Array) Insert(index uint, v int) error {
	if a.Len() == uint(cap(a.values)) {
		return errors.New("Array is full!")
	}
	if index != a.length && a.isIndexOutOfRange(index) {
		return errors.New("Index out of range!")
	}
	for i := a.length; i > index; i-- {
		a.values[i] = a.values[i-1]
	}
	a.values[index] = v
	a.length++
	return nil
}

// InsertToTail inserts v on tail
func (a *Array) InsertToTail(v int) error {
	return a.Insert(a.Len(), v)
}

// Delete deletes element at index
func (a *Array) Delete(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("Index out of range!")
	}
	v := a.values[index]
	for i := index; i < a.Len()-1; i++ {
		a.values[i] = a.values[i+1]
	}
	a.length--
	return v, nil
}

// Print prints the array
func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.values[i])
	}
	fmt.Println(format)
}
