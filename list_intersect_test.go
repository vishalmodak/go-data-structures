package main

import (
	"./ds"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	var intersectionTests = []struct {
		list1    *ds.LinkedList
		list2    *ds.LinkedList
		expected string
	}{
		{buildList(3, 6, 9, 11, 12, 15), buildList(6, 5, 7, 11, 12, 15), "6,11,12,15"},
		{buildList(3, 6, 9, 11), buildList(11, 12, 15), "11"},
		{buildList(3, 6, 9, 11), buildList(1, 2, 5), ""},
	}

	for _, tt := range intersectionTests {
		intersectionList := FindIntersection(tt.list1, tt.list2)
		if intersectionList.String() != tt.expected {
			t.Errorf("Intersection list incorrect, got %s, expected %s", intersectionList, tt.expected)
		}
	}

}

func buildList(values ...int) *ds.LinkedList {
	list := new(ds.LinkedList)
	for _, v := range values {
		list.Insert(v)
	}
	return list
}
