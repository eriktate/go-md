package md

import "fmt"

//List is a doubly linked list that keeps track of its own length.
type List struct {
	head   *ListItem
	tail   *ListItem
	length int
}

//ListItem is a single node in a List. It keeps track of both of its neighbors as well as some type of data.
type ListItem struct {
	prev *ListItem
	next *ListItem
	data interface{}
}

//NewList returns a pointer to an empty List struct. If you need a List with an element in it to start, then you can
//call Push with your data to accomplish that.
func NewList() *List {
	return &List{head: nil, tail: nil, length: 0}
}

//NewListItem returns a pointer to a valid ListItem. Nils can be passed into the prev/next pointers if the item does not
//know its neighbors.
func NewListItem(prev, next *ListItem, data interface{}) *ListItem {
	return &ListItem{prev: prev, next: next, data: data}
}

//Get attempts to retrieve the item at a specific index. If the index is out of bounds, a nil pointer and an error will
//be returned.
func (l *List) Get(index int) (*ListItem, error) {
	item := l.head
	for i := 0; i < l.length; i++ {
		if i == index {
			return item, nil
		}

		if item.next != nil {
			item = item.next
		} else {
			return nil, fmt.Errorf("Failed to Get item at index: %d. List is length: %d", index, l.length)
		}
	}
	return nil, fmt.Errorf("Failed to get item at index: %d. List is length: %d", index, l.length)
}

//Push adds an item to the tail of the list, replacing head and tail if the list is empty.
func (l *List) Push(data interface{}) {
	if l.head == nil {
		newItem := NewListItem(nil, nil, data)
		l.head = newItem
		l.tail = newItem
	} else {
		newItem := NewListItem(l.tail, nil, data)
		l.tail.next = newItem
	}
	l.length++
}

//Insert attempts to insert an item at an index.
func (l *List) Insert(index int, data interface{}) error {
	oldItem, err := l.Get(index)

	if err != nil {
		return err
	}

	newItem := NewListItem(oldItem.prev, oldItem, data)
	oldItem.prev.next = newItem
	oldItem.prev = newItem
	return nil
}

//InsertAfter inserts an item after a given node.
func (l *List) InsertAfter(node *ListItem, data interface{}) {
	newItem := NewListItem(node, node.next, data)
	node.next = newItem
}

//InsertBefore inserts an item before a given node.
func (l *List) InsertBefore(node *ListItem, data interface{}) {
	newItem := NewListItem(node.prev, node, data)
	node.prev.next = newItem
	node.prev = newItem
}

//RemoveAt attempts to remove an item at a specific index. It returns the removed item.
func (l *List) RemoveAt(index int) (*ListItem, error) {
	removed, err := l.Get(index)

	if err != nil {
		return nil, err
	}

	return l.Remove(removed), nil
}

//Remove will remove the links to a given ListItem and return that ListItem.
func (l *List) Remove(node *ListItem) *ListItem {
	if node == l.head {
		if node.next == nil {
			l.head = nil
			l.tail = nil
			l.length = 0
		} else {
			l.head = node.next
			node.next = nil
		}
		return node
	}

	if node == l.tail {
		if node.prev == nil {
			l.head = nil
			l.tail = nil
			l.length = 0
		} else {
			l.tail = node.prev
			node.prev = nil
		}
		return node
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil
	node.next = nil

	return node
}
