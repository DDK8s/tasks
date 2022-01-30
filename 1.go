package main

import (
	"fmt"
)

type elem struct {
	name string
	next *elem
}

type singleList struct {
	len  int
	head *elem
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) AddBack(name string) *elem{
	elem := &elem{
		name: name,
	}
	if s.head == nil {
		s.head = elem
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = elem
	}
	s.len++
	return elem
}

func (s *singleList) showList() error {
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}

	return nil
}


func main() {
	singleList := initList()

	singleList.AddBack("A")
	singleList.AddBack("B")
	singleList.AddBack("C")
	singleList.AddBack("D")
	singleList.AddBack("E")

	fast := singleList.head
	slow := singleList.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		//Если коллизия найдена.
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.next == nil {
		fmt.Println("No loop.")
		return
	}


	slow = singleList.head
	for slow != fast {
		slow = slow.next
		fast = singleList.head.next
	}
	fmt.Println("Loop starting point is ", fast.name)

	//Вывести весь список.
	/*err := singleList.showList()
	if err != nil {
		fmt.Println(err.Error())
	}*/

}
