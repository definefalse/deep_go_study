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
	if m.root == nil {
		m.root = &Node{key, value, nil, nil}
		m.size++
		return
	}

	cur := m.root
	for {
		if key < cur.key {
			if cur.left == nil {
				cur.left = &Node{key, value, nil, nil}
				m.size++
				return
			}
			cur = cur.left
			continue
		}
		if key > cur.key {
			if cur.right == nil {
				cur.right = &Node{key, value, nil, nil}
				m.size++
				return
			}
			cur = cur.right
			continue
		}
		cur.value = value
		return
	}
}

func (m *OrderedMap) Erase(key int) {
	var parent *Node
	cur := m.root

	for cur != nil && cur.key != key {
		parent = cur
		if key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}

	if cur == nil {
		fmt.Printf("Key %v not found\n", key)
		return
	}

	if cur.left != nil && cur.right != nil {
		parentChild := cur
		child := cur.right

		for child.left != nil {
			parentChild = child
			child = child.left
		}

		cur.key = child.key
		cur.value = child.value

		parent = parentChild
		cur = child
	}

	var child *Node
	if cur.left != nil {
		child = cur.left
	} else {
		child = cur.right
	}

	if parent == nil {
		m.root = child
	} else if parent.left == cur {
		parent.left = child
	} else {
		parent.right = child
	}

	m.size--
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
