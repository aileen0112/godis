package list

// List includes root element and length of list
type List struct {
	Root Element
	Len  int
}

// Element is the list's elemennt
type Element struct {
	next  *Element
	prev  *Element
	list  *List
	value interface{}
}

// Create return a new list
func Create() *List {
	return new(List)
}

/*
func Init() *List {
}
*/

// AddHead insert value v to the head of list
func (l *List) AddHead(v interface{}) *Element {
	return l.insert(&Element{value: v}, &l.Root)
}

//func AddTail(node Node) error {}
func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.next = n
	e.prev = at
	n.prev = e
	e.list = l
	l.Len++
	return e
}

// Len return the length of list
func Len(l *List) int {
	return l.Len
}

// First return the root element of list
func (l *List) First() *Element {
	return &l.Root
}

//func Last(list List) (Node, error) {}
