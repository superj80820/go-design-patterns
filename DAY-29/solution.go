package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type IterableSliceString []string

func (i IterableSliceString) Iterator() Iterator {
	return &SliceStringIterator{
		original: i,
		index:    0,
	}
}

type SliceStringIterator struct {
	original IterableSliceString
	index    int
}

func (s *SliceStringIterator) HasNext() bool {
	return s.index < len(s.original)
}

func (s *SliceStringIterator) Next() interface{} {
	item := s.original[s.index]
	s.index++
	return item
}

type IterableString string

func (i IterableString) Iterator() Iterator {
	return &StringIterator{
		original: i,
		index:    0,
	}
}

type StringIterator struct {
	original IterableString
	index    int
}

func (s *StringIterator) HasNext() bool {
	return s.index < len(s.original)
}

func (s *StringIterator) Next() interface{} {
	item := string(s.original[s.index])
	s.index++
	return item
}

func PrintAllItems(iterator Iterator) {
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}

func main() {
	PrintAllItems(IterableSliceString{"a", "b", "c"}.Iterator())
	PrintAllItems(IterableString("abcd").Iterator())
}
