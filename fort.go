package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ele struct {
	val  int
	next *ele
}

type list struct {
	len  int
	head *ele
}

func initList() *list {
	return &list{}
}

func (s *list) push(val int) {
	ele := &ele{
		val: val,
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
}

func (s *list) pop() {
	if s.head != nil {
		var prev *ele
		current := s.head
		for current.next != nil {
			prev = current
			current = current.next
		}
		if prev != nil {
			prev.next = nil
		} else {
			s.head = nil
		}
		s.len--
	}
}
func printList(l list) {
	current := l.head
	if l.len == 0 {
		fmt.Println("No nodes in list")
	} else {

		for current.next != nil {
			fmt.Print(current.val, " ")
			current = current.next
		}
		fmt.Print(current.val, " ")
	}
}
func (s *list) Size() int {
	return s.len
}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func main() {
	extra_tokens := make(map[string][]string)
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("<Ф ")
		str, _ := reader.ReadString('\n')
		str = strings.Trim(str, "\r\n")
		slice := strings.Split(str, " ")
		for index, element := range slice {
			switch element {
			case ":":
				i := Find(slice, ";")
				fmt.Println(slice)
				fmt.Println(i)
				if i != -1 {
					extra_tokens[slice[index+1]] = slice[index+1 : i]
				} else {
					fmt.Println("Не удалось найти ';'")
				}
			}
		}
		for key, value := range extra_tokens {
			fmt.Println(key, value)
		}
	}
}
