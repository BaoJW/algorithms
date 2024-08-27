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
