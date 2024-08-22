package hashMap

import (
	"errors"
	"fmt"
	"reflect"
)

// 拉链法实现hash

type KVNode[K comparable, V any] struct {
	key   K
	value V
}

type MyChainingHashMap[K comparable, V any] struct {
	table [][]KVNode[K, V]
	size  int
	cap   int
}

const INIT_CAP = 4

func NewMyChainingHashMap[K comparable, V any]() *MyChainingHashMap[K, V] {
	return NewMyChainingHashMapWithCapacity[K, V](INIT_CAP)
}

func NewMyChainingHashMapWithCapacity[K comparable, V any](initCapacity int) *MyChainingHashMap[K, V] {
	if initCapacity < 1 {
		initCapacity = 1
	}
	table := make([][]KVNode[K, V], initCapacity)
	return &MyChainingHashMap[K, V]{table: table, size: 0, cap: initCapacity}
}

func (m *MyChainingHashMap[K, V]) Put(key K, value V) error {
	if key == nil {
		return errors.New("key is nil")
	}
	index := m.hash(key)
	list := m.table[index]

	for i, node := range list {
		if reflect.DeepEqual(node.key, key) {
			list[i].value = value
			return nil
		}
	}

	m.table[index] = append(m.table[index], KVNode[K, V]{key: key, value: value})
	m.size++

	if float64(m.size) >= float64(m.cap)*0.75 {
		m.resize(m.cap * 2)
	}
	return nil
}

func (m *MyChainingHashMap[K, V]) Remove(key K) error {
	if key == nil {
		return errors.New("key is nil")
	}
	index := m.hash(key)
	list := m.table[index]

	for i, node := range list {
		if reflect.DeepEqual(node.key, key) {
			m.table[index] = append(list[:i], list[i+1:]...) // Remove node
			m.size--

			if float64(m.size) <= float64(m.cap)/8 {
				m.resize(m.cap / 4)
			}
			return nil
		}
	}
	return nil
}

func (m *MyChainingHashMap[K, V]) Get(key K) (V, error) {
	if key == nil {
		return *new(V), errors.New("key is nil")
	}
	index := m.hash(key)
	list := m.table[index]

	for _, node := range list {
		if reflect.DeepEqual(node.key, key) {
			return node.value, nil
		}
	}
	return *new(V), nil // Return zero value for V
}

func (m *MyChainingHashMap[K, V]) Keys() []K {
	keys := []K{}
	for _, list := range m.table {
		for _, node := range list {
			keys = append(keys, node.key)
		}
	}
	return keys
}

func (m *MyChainingHashMap[K, V]) Size() int {
	return m.size
}

func (m *MyChainingHashMap[K, V]) hash(key K) int {
	return int(hashCode(key)&0x7fffffff) % m.cap
}

func (m *MyChainingHashMap[K, V]) resize(newCap int) {
	newCap = max(newCap, 1)
	newMap := NewMyChainingHashMapWithCapacity[K, V](newCap)

	for _, list := range m.table {
		for _, node := range list {
			newMap.Put(node.key, node.value)
		}
	}
	m.table = newMap.table
	m.cap = newCap
}

func hashCode[K comparable](key K) int {
	return int(fmt.Sprintf("%v", key)[0]) // Simplified hash code
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
