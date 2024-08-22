package hashMap

import (
	"errors"
)

// 线性探查法实现哈希表
// 难点1：哈希表底层的table数组是一个一维数组，当发生哈希冲突时，我们需要往后找一个空位，如果一直找到数组的末尾还没找到空位，怎么办，这时候就需要转到数组的头部继续找空位，直到直到为止(环形数组技巧)
// 难点2: 删除操作较复杂，因为要维护元素的连续性 解决方案1: 数据搬移避免空洞；解决方案2：占位符标记删除，优点是不需要进行数据搬移，删除操作处理起来比较简单，缺点是不断地插入和删除元素，table 数组中会出现很多这样占位符，这样会增加连续元素的长度，进而降低 get 方法线性探查的效率

// MyLinearProbingHashMap rehash版本
type MyLinearProbingHashMap[K comparable, V any] struct {
	table []KVNode[K, V]
	size  int
	cap   int
}

func NewMyLinearProbingHashMap[K comparable, V any]() *MyLinearProbingHashMap[K, V] {
	return NewMyLinearProbingHashMapWithCapacity[K, V](INIT_CAP)
}

func NewMyLinearProbingHashMapWithCapacity[K comparable, V any](initCapacity int) *MyLinearProbingHashMap[K, V] {
	table := make([]KVNode[K, V], initCapacity)
	return &MyLinearProbingHashMap[K, V]{table: table, size: 0, cap: initCapacity}
}

func (m *MyLinearProbingHashMap[K, V]) Put(key K, val V) error {
	if key == nil {
		return errors.New("key is nil")
	}
	if m.size >= m.cap*3/4 {
		m.resize(m.cap * 2)
	}

	index := m.getKeyIndex(key)
	if m.table[index].key != nil {
		m.table[index].value = val
		return nil
	}

	m.table[index] = KVNode[K, V]{key: key, val: val}
	m.size++
	return nil
}

func (m *MyLinearProbingHashMap[K, V]) Remove(key K) error {
	if key == nil {
		return errors.New("key is nil")
	}
	if m.size <= m.cap/8 {
		m.resize(m.cap / 2)
	}

	index := m.getKeyIndex(key)
	if m.table[index].key == nil {
		return nil // key 不存在
	}

	m.table[index] = KVNode[K, V]{}
	m.size--

	// 进行 rehash
	index = (index + 1) % m.cap
	for m.table[index].key != nil {
		entry := m.table[index]
		m.table[index] = KVNode[K, V]{} // 清空当前节点
		m.size--
		m.Put(entry.key, entry.val) // 重新插入
	}
	return nil
}

func (m *MyLinearProbingHashMap[K, V]) Get(key K) (V, error) {
	if key == nil {
		return *new(V), errors.New("key is nil")
	}
	index := m.getKeyIndex(key)
	if m.table[index].key == nil {
		return *new(V), nil // key 不存在
	}
	return m.table[index].val, nil
}

func (m *MyLinearProbingHashMap[K, V]) Size() int {
	return m.size
}

func (m *MyLinearProbingHashMap[K, V]) getKeyIndex(key K) int {
	index := m.hash(key)
	for m.table[index].key != nil {
		if m.table[index].key == key {
			return index
		}
		index = (index + 1) % m.cap
	}
	return index
}

func (m *MyLinearProbingHashMap[K, V]) hash(key K) int {
	return int(key.(int)&0x7fffffff) % m.cap
}

func (m *MyLinearProbingHashMap[K, V]) resize(newCap int) {
	newMap := NewMyLinearProbingHashMapWithCapacity[K, V](newCap)
	for _, entry := range m.table {
		if entry.key != nil {
			newMap.Put(entry.key, entry.val)
		}
	}
	m.table = newMap.table
	m.cap = newCap
}

// MyLinearProbingHashMap2 特殊标记版本
type MyLinearProbingHashMap2[K comparable, V any] struct {
	table []KVNode[K, V]
	size  int
	cap   int
	DUMMY KVNode[K, V]
}

func NewMyLinearProbingHashMap2[K comparable, V any]() *MyLinearProbingHashMap2[K, V] {
	return NewMyLinearProbingHashMap2WithCapacity[K, V](INIT_CAP)
}

func NewMyLinearProbingHashMap2WithCapacity[K comparable, V any](cap int) *MyLinearProbingHashMap2[K, V] {
	table := make([]KVNode[K, V], cap)
	return &MyLinearProbingHashMap2[K, V]{table: table, size: 0, cap: cap, DUMMY: KVNode[K, V]{}}
}

// 添加 key -> val 键值对
func (m *MyLinearProbingHashMap2[K, V]) Put(key K, val V) error {
	if key == nil {
		return errors.New("key is nil")
	}

	if m.size >= m.cap*3/4 {
		m.resize(m.cap * 2)
	}

	index := m.getKeyIndex(key)
	if index != -1 {
		// key 已存在，修改对应的 val
		m.table[index].val = val
		return nil
	}

	// key 不存在
	x := KVNode[K, V]{key: key, val: val}
	index = m.hash(key)
	for m.table[index].key != nil && m.table[index] != m.DUMMY {
		index = (index + 1) % m.cap
	}
	m.table[index] = x
	m.size++
	return nil
}

// 删除 key 和对应的 val，并返回 val
func (m *MyLinearProbingHashMap2[K, V]) Remove(key K) error {
	if key == nil {
		return errors.New("key is nil")
	}

	if m.size < m.cap/8 {
		m.resize(m.cap / 2)
	}

	index := m.getKeyIndex(key)
	if index == -1 {
		return nil // key 不存在
	}

	m.table[index] = m.DUMMY
	m.size--
	return nil
}

// 返回 key 对应的 val
func (m *MyLinearProbingHashMap2[K, V]) Get(key K) (V, error) {
	if key == nil {
		return *new(V), errors.New("key is nil")
	}

	index := m.getKeyIndex(key)
	if index == -1 {
		return *new(V), nil // key 不存在
	}

	return m.table[index].val, nil
}

func (m *MyLinearProbingHashMap2[K, V]) Size() int {
	return m.size
}

// 对 key 进行线性探查，返回一个索引
func (m *MyLinearProbingHashMap2[K, V]) getKeyIndex(key K) int {
	step := 0
	for i := m.hash(key); m.table[i].key != nil; i = (i + 1) % m.cap {
		entry := m.table[i]
		if entry == m.DUMMY {
			continue
		}
		if entry.key == key {
			return i
		}
		step++
		if step == m.cap {
			m.resize(m.cap) // 清理占位符
			return -1
		}
	}
	return -1
}

// 哈希函数
func (m *MyLinearProbingHashMap2[K, V]) hash(key K) int {
	return (key.(int) & 0x7fffffff) % m.cap
}

// 扩容或缩容
func (m *MyLinearProbingHashMap2[K, V]) resize(cap int) {
	newMap := NewMyLinearProbingHashMap2WithCapacity[K, V](cap)
	for _, entry := range m.table {
		if entry.key != nil && entry != m.DUMMY {
			newMap.Put(entry.key, entry.val)
		}
	}
	m.table = newMap.table
	m.cap = cap
}
