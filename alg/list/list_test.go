package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 假设 SingleList 和 MergeList 已经定义并实现

func TestMergeList(t *testing.T) {
	tests := []struct {
		listAElements []int
		listBElements []int
		expected      []int
	}{
		{
			listAElements: []int{},
			listBElements: []int{0},
			expected:      []int{0},
		},
		{
			listAElements: []int{1, 2, 4},
			listBElements: []int{1, 3, 4},
			expected:      []int{1, 1, 2, 3, 4, 4},
		},
	}

	for _, tt := range tests {
		listA := NewSingleList()
		for _, val := range tt.listAElements {
			listA.AddElementToLast(val)
		}

		listB := NewSingleList()
		for _, val := range tt.listBElements {
			listB.AddElementToLast(val)
		}

		newList := MergeList(listA, listB)
		point := newList.head.Next

		// 提取结果
		var result []int
		for point != nil {
			result = append(result, point.Val)
			point = point.Next
		}

		// 验证结果
		assert.Equal(t, tt.expected, result)
	}
}

func TestMergeListV2(t *testing.T) {
	tests := []struct {
		listAElements []int
		listBElements []int
		expected      []int
	}{
		{
			listAElements: []int{},
			listBElements: []int{0},
			expected:      []int{0},
		},
		{
			listAElements: []int{1, 2, 4},
			listBElements: []int{1, 3, 4},
			expected:      []int{1, 1, 2, 3, 4, 4},
		},
		{
			listAElements: []int{2, 3, 5},
			listBElements: []int{1, 4, 6},
			expected:      []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		listA := createList(tt.listAElements)
		listB := createList(tt.listBElements)

		newList := MergeListV2(listA, listB)
		result := extractValues(newList)

		assert.Equal(t, tt.expected, result)
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		listAElements []int
		xValue        int
		expected      []int
	}{
		{
			listAElements: []int{1, 4, 3, 2, 5, 2},
			xValue:        3,
			expected:      []int{1, 2, 2, 4, 3, 5},
		},
		{
			listAElements: []int{2, 1},
			xValue:        2,
			expected:      []int{1, 2},
		},
	}

	for _, tt := range tests {
		listA := createList(tt.listAElements)

		newList := Partition(listA, tt.xValue)
		result := extractValues(newList)

		assert.Equal(t, tt.expected, result)
	}
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		listElements [][]int
		expected     []int
	}{
		{
			listElements: [][]int{
				{1, 4, 5},
				{2, 6},
				{1, 3, 4},
			},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		listk := make([]*Node, 0, len(tt.listElements))
		for _, elements := range tt.listElements {
			listk = append(listk, createList(elements))
		}

		newList := MergeKLists(listk)
		result := extractValues(newList)

		assert.Equal(t, tt.expected, result)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		listElements []int
		n            int
		expected     []int
	}{
		{
			listElements: []int{1, 2, 3, 4, 5},
			n:            2,
			expected:     []int{1, 2, 3, 5},
		},
		{
			listElements: []int{1},
			n:            1,
			expected:     nil,
		},
		{
			listElements: []int{1, 2},
			n:            1,
			expected:     []int{1},
		},
		{
			listElements: []int{1, 2, 3, 4, 5},
			n:            5,
			expected:     []int{2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		list := createList(tt.listElements)

		newList := removeNthFromEnd(list, tt.n)
		result := extractValues(newList)

		assert.Equal(t, tt.expected, result)
	}
}
