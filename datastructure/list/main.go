package main

import (
	"container/list"
	"fmt"
)

func main() {

	val2 := 18
	var intList list.List
	intList.PushBack(7)
	intList.PushBack(val2)
	intList.PushBack(35)

	printForward(intList)
	printBackWard(intList)

	elem := intList.Front()

	intList.InsertAfter(4, elem)

	printForward(intList)

	fmt.Println("Len : ", intList.Len())

}

func printForward(valList list.List) {
	for element := valList.Front(); element != nil; element=element.Next() {
		fmt.Println(element.Value.(int))
	}
}

func printBackWard(valList list.List) {
	for element := valList.Back(); element != nil; element=element.Prev() {
		fmt.Println(element.Value.(int))
	}
}
