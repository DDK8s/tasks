package main

import (
	"fmt"
)

type elem struct {
	name string
	next *ele
}

type sinList struct {
	len  int
	head *ele
}

func initList() *sinList {
	return &sinList{}
}

func (s *sinList) AddBack(name string) {
	ele := &ele{
		name: name,
	}
	if s.head == nil {
		s.head = ele
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = ele
	}
	s.len++
	return
}

func (s *sinList) showList() error {
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}

func main() {
	sinList := initList()

	sinList.AddBack("A")
	sinList.AddBack("B")
	sinList.AddBack("C")
	sinList.AddBack("D")
	sinList.AddBack("E")
	sinList.AddBack("C")

	fmt.Print(sinList.len)

	err := sinList.showList()
	if err != nil {
		fmt.Println(err.Error())
	}


}
