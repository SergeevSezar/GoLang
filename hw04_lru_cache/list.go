package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: nil, Prev: nil}

	if l.front == nil {
		l.front = newItem
		l.back = newItem
	} else {
		oldFront := l.front
		newItem.Next = oldFront
		oldFront.Prev = newItem
		l.front = newItem
	}
	l.len++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: nil, Prev: nil}

	if l.back == nil {
		l.back = newItem
		l.front = newItem
	} else {
		oldBack := l.back
		newItem.Prev = oldBack
		oldBack.Next = newItem
		l.back = newItem
	}
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	prevItem := i.Prev
	nextItem := i.Next

	if prevItem != nil {
		prevItem.Next = nextItem
	} else {
		l.front = nextItem
	}

	if nextItem != nil {
		nextItem.Prev = prevItem
	} else {
		l.back = prevItem
	}

	i.Next = nil
	i.Prev = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front == i {
		return
	}

	prevItem := i.Prev
	nextItem := i.Next

	if prevItem != nil {
		prevItem.Next = nextItem
	} else {
		l.front = nextItem
	}

	if nextItem != nil {
		nextItem.Prev = prevItem
	} else {
		l.back = prevItem
	}

	i.Prev = nil
	i.Next = l.front

	if l.front != nil {
		l.front.Prev = i
	}

	l.front = i

	if l.back == nil {
		l.back = i
	}
}
