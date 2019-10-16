package linked_list

// copy and learn from go/1.12.6/libexec/src/container/list/list.go
// ![source-code]
type Element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List represents a doubly linked list.
type List struct {
	root Element
	len  int
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// !! e must not be nil...
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// PushFront, PushBack, InsertBefore, InsertAfter insert new element
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// MoveToFront
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, l.root.prev)
}

// MoveBefore moves e to its new position before mark
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves e to its new position after mark
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
	return e
}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

// move moves e to next to at and returns e.
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	// remove e
	e.prev.next = e.next
	e.next.prev = e.prev

	// insert e next to at
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e

	return e
}
