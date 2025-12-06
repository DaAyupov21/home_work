package hw04lrucache

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type list struct {
	len         int
	front, back *ListItem
}

func NewList() List { return &list{} }

func (l *list) Len() int         { return l.len }
func (l *list) Front() *ListItem { return l.front }
func (l *list) Back() *ListItem  { return l.back }

func (l *list) PushFront(v interface{}) *ListItem {
	it := &ListItem{Value: v, Next: l.front}
	if l.front != nil {
		l.front.Prev = it
	} else {
		l.back = it // список был пуст
	}
	l.front = it
	l.len++
	return it
}

func (l *list) PushBack(v interface{}) *ListItem {
	it := &ListItem{Value: v, Prev: l.back}
	if l.back != nil {
		l.back.Next = it
	} else {
		l.front = it // список был пуст
	}
	l.back = it
	l.len++
	return it
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next // удаляем голову
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev // удаляем хвост
	}
	i.Prev, i.Next = nil, nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || i == l.front {
		return
	}
	// вынимаем i
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev // i был хвостом
	}
	// вставляем в голову
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
