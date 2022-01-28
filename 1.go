package main

import (
	"fmt"
)

type elem struct {
	name string
	next *elem
}

type sinList struct {
	len  int
	head *elem
}

func initList() *sinList {
	return &sinList{}
}

func (s *sinList) AddBack(name string) string{
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
	return name
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

func (s *sinList) asd() error {
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
	
	err := sinList.showList()
	if err != nil {
		fmt.Println(err.Error())
	}


}
