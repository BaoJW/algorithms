package array

import "errors"

// 环形数组
// 其实本质上是一种技巧，让我们用O(1)的时间在数组头部增删元素
// 首先从物理层面数组不可能是环形的，因为数组就是一块线性连续的内存空间，不会有环的概念
// 我们说的环形数组，其实是逻辑层面把数组变成了环形
// 核心原理就是它维护了两个指针start, end，start指向第一个有效元素的索引，end指向最后一个有效元素的下一个位置的索引(左闭右开模式)
// 当我们在数组头部添加或删除元素时，只需要移动start索引，而在数组尾部添加或删除元素时，只需要移动end索引
// 当start,end移动超出数组边界(<0或>=arr.length)时，我们可以通过求模运算%让它们转一圈到数组头部或尾部继续工作，这样就实现了环形数组的效果

// 引发的思考：我们知道数组增删头部元素的效率是O(N)，因为需要数据搬移，但是环形数组通过取模计算索引其实是能实现O(1)时间复杂度的效率的，但为何编程语言的标准库中实现动态数组并没有使用环形数组？
// 一个可能的答案：如果使用环形数组，增删改查的所有操作都会涉及%求模运算，这个操作是比较耗性能的，尤其像数组的 get 方法，调用频率会非常非常高，如果每次调用都多一步 % 运算，加起来的性能损耗远大于环形数组带来的收益，因为数组很少在头部增删元素。如果你非要在头部增删，应该使用更合适的其他数据结构

type CycleArray[E any] struct {
	arr   []E
	start int // 开始指针索引下标
	end   int // 结束指针索引下标
	count int // 数组内的元素个数
	size  int // 数组容量
}

func NewCycleArray[E any]() *CycleArray[E] {
	return NewCycleArrayWithSize[E](1)
}

func NewCycleArrayWithSize[E any](size int) *CycleArray[E] {
	return &CycleArray[E]{
		arr:   make([]E, size),
		start: 0,
		end:   0,
		count: 0,
		size:  size,
	}
}

// 自动扩缩容辅助函数
func (c *CycleArray[E]) resize(newSize int) {
	newArr := make([]E, newSize)
	for i := 0; i < c.count; i++ {
		newArr[i] = c.arr[(c.start+i)%c.size]
	}
	c.arr = newArr
	c.start = 0
	c.end = c.count
	c.size = newSize
}

// AddFirst 在数组头部添加元素，时间复杂度 O(1)
func (c *CycleArray[E]) AddFirst(val E) {
	if c.isFull() {
		c.resize(c.size * 2)
	}
	c.start = (c.start - 1 + c.size) % c.size
	c.arr[c.start] = val
	c.count++
}

// RemoveFirst 删除数组头部元素，时间复杂度 O(1)
func (c *CycleArray[E]) RemoveFirst() (E, error) {
	if c.isEmpty() {
		return nil, errors.New("array is empty")
	}
	removedValue := c.arr[c.start]
	c.arr[c.start] = *new(E) // 清空元素
	c.start = (c.start + 1) % c.size
	c.count--
	if c.count > 0 && c.count == c.size/4 {
		c.resize(c.size / 2)
	}
	return removedValue, nil
}

// AddLast 在数组尾部添加元素，时间复杂度 O(1)
func (c *CycleArray[E]) AddLast(val E) {
	if c.isFull() {
		c.resize(c.size * 2)
	}
	c.arr[c.end] = val
	c.end = (c.end + 1) % c.size
	c.count++
}

// RemoveLast 删除数组尾部元素，时间复杂度 O(1)
func (c *CycleArray[E]) RemoveLast() error {
	if c.isEmpty() {
		return errors.New("array is empty")
	}
	c.end = (c.end - 1 + c.size) % c.size
	c.arr[c.end] = *new(E) // 清空元素
	c.count--
	if c.count > 0 && c.count == c.size/4 {
		c.resize(c.size / 2)
	}
	return nil
}

// GetFirst 获取数组头部元素，时间复杂度 O(1)
func (c *CycleArray[E]) GetFirst() (E, error) {
	if c.isEmpty() {
		return *new(E), errors.New("array is empty")
	}
	return c.arr[c.start], nil
}

// GetLast 获取数组尾部元素，时间复杂度 O(1)
func (c *CycleArray[E]) GetLast() (E, error) {
	if c.isEmpty() {
		return *new(E), errors.New("array is empty")
	}
	return c.arr[(c.end-1+c.size)%c.size], nil
}

func (c *CycleArray[E]) isFull() bool {
	return c.count == c.size
}

func (c *CycleArray[E]) Size() int {
	return c.count
}

func (c *CycleArray[E]) isEmpty() bool {
	return c.count == 0
}
