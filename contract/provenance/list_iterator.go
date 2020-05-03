package main

type Iterator interface {
	Value() interface{}
	Next() bool
}

type iterator struct {
	counter uint64
	total uint64
	list List
}

func NewListIterator(l List) Iterator {
	return &iterator{
		counter: 0,
		total: l.Length(),
		list: l,
	}
}

func (i *iterator) Value() interface{} {
	index := i.counter
	i.counter++
	return i.list.Get(index)
}

func (i *iterator) Next() bool {
	return i.counter < i.total
}