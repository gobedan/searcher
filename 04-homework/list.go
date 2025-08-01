// Реализуация двусвязного списка вместе с базовыми операциями.
package list

import (
	"fmt"
)

// List - двусвязный список.
type List struct {
	root *Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New создаёт список и возвращает указатель на него.
func New() *List {
	var l List
	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = &e
	if e.next != l.root {
		e.next.prev = &e
	}
	return &e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *List {
	l.root.next = l.root.next.next
	l.root.next.prev = l.root
	return l
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	el := l.root.next
	straight := true
	for el.next != l.root {
		el = el.next
		// чтобы начать разворачивать элементы только на обратном пути
		if !straight {
			el.next, el.prev = el.prev, el.next
		}
		// дойти до конца и начать обход в обратку от хвоста
		if el.next == l.root && straight {
			l.root.next = el
			el.next, el.prev = el.prev, el.next
			straight = false
		}
	}

	return l
}
