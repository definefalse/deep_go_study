package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

type OrderedMap struct {
	root *Node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	var inserted bool
	m.root, inserted = insertRecursive(m.root, key, value)
	if inserted {
		m.size++
	}
}

func insertRecursive(n *Node, key, value int) (*Node, bool) {
	if n == nil {
		return &Node{key: key, value: value}, true
	}

	if key < n.key {
		var inserted bool
		n.left, inserted = insertRecursive(n.left, key, value)
		return n, inserted
	}

	if key > n.key {
		var inserted bool
		n.right, inserted = insertRecursive(n.right, key, value)
		return n, inserted
	}

	n.value = value
	return n, false
}

func (m *OrderedMap) Erase(key int) {
	var erased bool
	m.root, erased = erase(m.root, key)
	if erased {
		m.size--
	}
}

func erase(n *Node, key int) (*Node, bool) {
	if n == nil {
		return nil, false
	}

	if key < n.key {
		var erased bool
		n.left, erased = erase(n.left, key)
		return n, erased
	}

	if key > n.key {
		var erased bool
		n.right, erased = erase(n.right, key)
		return n, erased
	}

	if n.left == nil {
		return n.right, true
	}
	if n.right == nil {
		return n.left, true
	}

	successor := minNode(n.right)
	n.key = successor.key
	n.value = successor.value
	n.right, _ = erase(n.right, successor.key)
	return n, true
}

func minNode(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (m *OrderedMap) Contains(key int) bool {
	cur := m.root
	for cur != nil {
		if cur.key == key {
			return true
		}
		if cur.key > key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return false
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	stack := make([]*Node, 0)
	cur := m.root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		action(cur.key, cur.value)

		cur = cur.right
	}
}

func TestOrderedMap(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())

	data.Insert(14, 10)
	assert.Equal(t, 7, data.Size())

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(4)
	data.Erase(5)
	data.Erase(10)
	data.Erase(12)
	data.Erase(30)

	var lastKeys []int
	data.ForEach(func(key, _ int) {
		lastKeys = append(lastKeys, key)
	})
	fmt.Print(lastKeys)
	assert.Equal(t, 0, data.Size())
}
