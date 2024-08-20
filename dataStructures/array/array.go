package array

import (
	"errors"
	"fmt"
)

// 数组(顺序存储)  静态数组  动态数组
// 静态数组：一块连续的内存空间，我们可以通过索引来访问这块内存空间中的元素，这是数组的原始形态
// 动态数组：是编程语言方便我们使用，静态数组的基础上帮我们添加了一些常用的 API，比如增删改查等，这些 API 可以让我们更方便地操作数组元素，不用自己去写代码实现这些操作

// 静态数组中的增删改查
// 1.增
// 1.1 原始数组容量未满的情况下，数组末尾增加(append)元素
// 时间复杂度：O(1)
func append() {
	// 大小为 10 的数组已经装了 4 个元素
	arr := make([]int, 10)
	for i := 0; i < 4; i++ {
		arr[i] = i
	}

	// 现在想在数组末尾追加一个元素 4
	arr[4] = 4

	// 再在数组末尾追加一个元素 5
	arr[5] = 5

	// 依此类推
	// ...

	// 打印数组
	fmt.Println(arr)
}

// 1.2 原始数组容量未满的情况下，数组中间插入(insert)元素
// 流程上是：先进行数据搬移，给新元素腾出空位，然后插入新元素
// 时间复杂度：O(N)
func insert() {
	// 大小为 10 的数组已经装了 4 个元素
	arr := make([]int, 10)
	for i := 0; i < 4; i++ {
		arr[i] = i
	}

	// 在第 3 个位置插入元素 666
	// 需要把第 3 个位置及之后的元素都往后移动一位
	for i := 4; i > 2; i-- {
		arr[i] = arr[i-1]
	}

	// 现在第 3 个位置空出来了，可以插入新元素
	arr[2] = 666

	// 打印数组
	fmt.Println(arr)
}

// 1.3 原始数组容量已满的情况下，数组增加元素操作
// 静态数组是在初始化时就已经分配好了一块连续的内存，无法在容量已满之后在其后面直接追加内存，因为你这块连续内存后面的内存空间可能已经被其他程序占用了，不能说你想要就给你
// 所以想要在此时对数组进行增加元素的操作，只能进行扩容
// 时间复杂度: O(N)
func scaling() {
	// 大小为 10 的数组已经装满了
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	// 现在想在数组末尾追加一个元素 10
	// 需要先扩容数组
	newArr := make([]int, 20)

	// 把原来的 10 个元素复制过去
	copy(newArr, arr)

	// 在新的大数组中追加新元素
	newArr[10] = 10

	// 打印新的数组
	fmt.Println(newArr)
}

// 2. 删
// 2.1 删除末尾元素
// 删除数组尾部元素的本质就是进行一次随机访问
// 时间复杂度: O(1)
func deleteLastElement() {
	// 大小为 10 的数组已经装了 5 个元素
	arr := make([]int, 10)
	for i := 0; i < 5; i++ {
		arr[i] = i
	}

	// 删除末尾元素，暂时用 -1 代表元素已删除
	arr[4] = -1

	// 打印数组
	fmt.Println(arr)
}

// 2.2 删除中间元素
// 本质就是数据搬移，把被删元素后面的元素都往前移动一位，保持数组元素的连续性
// 时间复杂度: O(N)
func deleteIntermediateElements() {
	// 大小为 10 的数组已经装了 5 个元素
	arr := make([]int, 10)
	for i := 0; i < 5; i++ {
		arr[i] = i
	}

	// 删除 arr[1]
	// 需要把 arr[1] 之后的元素都往前移动一位
	for i := 1; i < 4; i++ {
		arr[i] = arr[i+1]
	}

	// 最后一个元素置为 -1 代表已删除
	arr[4] = -1

	// 打印数组
	fmt.Println(arr)
}

// 3 查
// 给定指定索引，查询索引对应的元素的值
// 时间复杂度 O(1)

// 4 改
// 给定指定索引，修改索引对应的元素的值
// 时间复杂度 O(1)

// 动态数组
// 关键点1: 自动扩缩容
// 定义一个简单的策略：
//	当数组元素个数达到底层静态数组的容量上限时，扩容为原来的2倍
//  当数组元素个数缩减到底层静态数组的容量的1/4时，缩容为原来的1/2
// 关键点2：索引越界的检查
// 	checkElementIndex  index < size
//  checkPositionIndex index == size
// 关键点3：删除元素谨防内存泄漏
// 垃圾回收机制是基于图算法的可达性分析，如果一个对象再也无法被访问到，那么这个对象占用的内存才会被释放；否则，垃圾回收器会认为这个对象还在使用中，就不会释放这个对象占用的内存。
// 如果你不执行 data[size - 1] = null 这行代码，那么 data[size - 1] 这个引用就会一直存在，你可以通过 data[size - 1] 访问这个对象，所以这个对象被认为是可达的，它的内存就一直不会被释放，进而造成内存泄漏。

// MyArrayList 动态数组的实现
type MyArrayList struct {
	data []interface{}
	size int
}

const INIT_CAP = 1

func NewMyArrayList() *MyArrayList {
	return NewMyArrayListWithCapacity(INIT_CAP)
}

func NewMyArrayListWithCapacity(initCapacity int) *MyArrayList {
	return &MyArrayList{
		data: make([]interface{}, initCapacity),
		size: 0,
	}
}

// 增
func (list *MyArrayList) AddLast(e interface{}) {
	cap := len(list.data)
	if list.size == cap {
		list.resize(2 * cap)
	}
	list.data[list.size] = e
	list.size++
}

func (list *MyArrayList) Add(index int, e interface{}) error {
	if index < 0 || index > list.size {
		return errors.New("index out of bounds")
	}

	cap := len(list.data)
	if list.size == cap {
		list.resize(2 * cap)
	}

	for i := list.size - 1; i >= index; i-- {
		list.data[i+1] = list.data[i]
	}
	list.data[index] = e
	list.size++
	return nil
}

func (list *MyArrayList) AddFirst(e interface{}) {
	list.Add(0, e)
}

// 删
func (list *MyArrayList) RemoveLast() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("no element to remove")
	}

	cap := len(list.data)
	if list.size == cap/4 {
		list.resize(cap / 2)
	}

	deletedVal := list.data[list.size-1]
	list.data[list.size-1] = nil // 防止内存泄漏
	list.size--

	return deletedVal, nil
}

func (list *MyArrayList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("index out of bounds")
	}

	cap := len(list.data)
	if list.size == cap/4 {
		list.resize(cap / 2)
	}

	deletedVal := list.data[index]
	for i := index + 1; i < list.size; i++ {
		list.data[i-1] = list.data[i]
	}
	list.data[list.size-1] = nil // 防止内存泄漏
	list.size--

	return deletedVal, nil
}

func (list *MyArrayList) RemoveFirst() (interface{}, error) {
	return list.Remove(0)
}

// 查
func (list *MyArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("index out of bounds")
	}
	return list.data[index], nil
}

// 改
func (list *MyArrayList) Set(index int, element interface{}) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("index out of bounds")
	}
	oldVal := list.data[index]
	list.data[index] = element
	return oldVal, nil
}

// 工具方法
func (list *MyArrayList) Size() int {
	return list.size
}

func (list *MyArrayList) IsEmpty() bool {
	return list.size == 0
}

// 将 data 的容量改为 newCap
func (list *MyArrayList) resize(newCap int) {
	temp := make([]interface{}, newCap)
	copy(temp, list.data)
	list.data = temp
}

func main() {
	// 初始容量设置为 3
	arr := NewMyArrayListWithCapacity(3)

	// 添加 5 个元素
	for i := 1; i <= 5; i++ {
		arr.AddLast(i)
	}

	arr.Remove(3)
	arr.Add(1, 9)
	arr.AddFirst(100)
	val, _ := arr.RemoveLast()

	for i := 0; i < arr.Size(); i++ {
		if v, err := arr.Get(i); err == nil {
			fmt.Println(v)
		}
	}

	fmt.Println("Removed value:", val)
}
