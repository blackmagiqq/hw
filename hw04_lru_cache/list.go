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
	countOfListItems int
	frontItem        *ListItem
	backItem         *ListItem
}

func (l *list) Len() int {
	return l.countOfListItems
}

func (l *list) Front() *ListItem {
	if l.frontItem != nil {
		return l.frontItem
	}
	return l.backItem
}

func (l *list) Back() *ListItem {
	if l.backItem != nil {
		return l.backItem
	}
	return l.frontItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	currentFrontItem := l.frontItem
	currentBackItem := l.backItem

	newItem := &ListItem{
		Value: v,
		Next:  currentFrontItem,
	}

	if currentFrontItem != nil {
		currentFrontItem.Prev = newItem
	}
	if currentBackItem == nil {
		l.backItem = newItem
	}

	l.frontItem = newItem
	l.countOfListItems++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	currentBackItem := l.Back()

	newItem := &ListItem{
		Value: v,
		Prev:  currentBackItem,
	}

	if currentBackItem != nil {
		currentBackItem.Next = newItem
	}

	l.backItem = newItem
	l.countOfListItems++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if l.countOfListItems == 0 {
		return
	}
	l.countOfListItems--

	prevListItem := i.Prev
	nextListItem := i.Next

	if nextListItem != nil && prevListItem != nil {
		prevListItem.Next = nextListItem
		nextListItem.Prev = prevListItem
		return
	}

	if nextListItem != nil {
		nextListItem.Prev = nil
		l.frontItem = nextListItem
		return
	}

	if prevListItem != nil {
		prevListItem.Next = nil
		l.backItem = prevListItem
		return
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if l.countOfListItems <= 1 {
		return
	}
	if l.frontItem == i {
		return
	}

	prevListItem := i.Prev
	nextListItem := i.Next

	if nextListItem != nil && prevListItem != nil {
		prevListItem.Next = nextListItem
		nextListItem.Prev = prevListItem

		i.Next = l.frontItem
		l.frontItem.Prev = i

		i.Prev = nil
		l.frontItem = i
		return
	}

	if prevListItem != nil {
		prevListItem.Next = nil

		i.Next = l.frontItem
		l.frontItem.Prev = i
		i.Prev = nil

		l.backItem = prevListItem
		l.frontItem = i
	}
}

func NewList() List {
	return new(list)
}
