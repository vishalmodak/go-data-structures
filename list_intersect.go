package main

import (
	"./ds"
)

func FindIntersection(list1 *ds.LinkedList, list2 *ds.LinkedList) *ds.LinkedList {
	intersectionList := new(ds.LinkedList)
	list1Idx := list1.Head()
	for list1Idx != nil {
		list2Idx := list2.Head()
		for list2Idx != nil {
			if list1Idx.GetData() == list2Idx.GetData() {
				intersectionList.Insert(list1Idx.GetData())
			}
			list2Idx = list2Idx.Next()
		}
		list1Idx = list1Idx.Next()
	}
	return intersectionList
}
