package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func (s *list) pop() int {
	var val int = s.head.val
	if s.head != nil {
		var prev *ele
		current := s.head
		for current.next != nil {
			prev = current
			current = current.next
			val = current.val
		}
		if prev != nil {
			prev.next = nil
		} else {
			s.head = nil
		}
		s.len--
	}
	return val
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

//func (s *list) Size() int {
//	return s.len
//}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}
func is_digit(elem string) bool {
	if (elem[0] == '-') || (elem[0] == '0') || (elem[0] == '1') || (elem[0] == '2') || (elem[0] == '3') || (elem[0] == '4') ||
		(elem[0] == '5') || (elem[0] == '6') || (elem[0] == '7') || (elem[0] == '8') || (elem[0] == '9') {
		return true
	}
	return false
}
func toInt(elem string) int {
	int2, _ := strconv.Atoi(elem)
	return int2
}
func split_into_array(str string) []string {
	word := ""
	var arr []string
	for _, char := range str {
		if char == ' ' {
			arr = append(arr, word)
			word = ""
		} else {
			word += string(char)
		}
	}
	arr = append(arr, word)
	return arr
}

func do_switch(arr []string, stack *list, m map[string][]string) {
	flag := false
	for _, elem := range arr {
		if flag == true {
			break
		}
		if is_digit(elem) {
			stack.push(toInt(elem))
		} else {
			switch elem {
			case ".":
				if stack.len > 0 {
					fmt.Println(stack.pop())
				} else {
					fmt.Println("no elems to destroy")
				}
			case "УДАЛ":
				if stack.len > 0 {
					stack.pop()
				} else {
					fmt.Println("no elems to destroy")
				}
			case "ДУБ":
				if stack.len > 0 {
					tmp := stack.pop()
					stack.push(tmp)
					stack.push(tmp)
				} else {
					fmt.Println("no elems to duplicate")
				}
			case "ОБМЕН":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_v)
					stack.push(tmp_u)
				} else {
					fmt.Println("no enough arguments to do swap")
				}
			case "+":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_v + tmp_u)
				} else {
					fmt.Println("no enough arguments to do +")
				}
			case "-":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_u - tmp_v)
				} else {
					fmt.Println("no enough arguments to do -")
				}
			case "*":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_u * tmp_v)
				} else {
					fmt.Println("no enough arguments to do *")
				}
			case "/":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_u / tmp_v)
				} else {
					fmt.Println("no enough arguments to do /")
				}
			case ":":
				flag = true
				i := Find(arr, ":")
				adding_new_keywords(i, stack, arr, m)
			default:
				if _, found := m[elem]; found {
					operate_with_map(elem, m, stack)
				} else {
					fmt.Println("unrecognized word")
				}
			}
		}
	}
}
func operate_with_map(elem string, m map[string][]string, stack *list) {
	for _, elem := range m[elem] {
		if is_digit(elem) {
			stack.push(toInt(elem))
		} else {
			switch elem {
			case ".":
				if stack.len > 0 {
					fmt.Println(stack.pop())
				} else {
					fmt.Println("no elems to destroy")
				}
			case "УДАЛ":
				if stack.len > 0 {
					stack.pop()
				} else {
					fmt.Println("no elems to destroy")
				}
			case "ДУБ":
				if stack.len > 0 {
					tmp := stack.pop()
					stack.push(tmp)
					stack.push(tmp)
				} else {
					fmt.Println("no elems to duplicate")
				}
			case "ОБМЕН":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_v)
					stack.push(tmp_u)
				} else {
					fmt.Println("no enough arguments to do swap")
				}
			case "+":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_v + tmp_u)
				} else {
					fmt.Println("no enough arguments to do +")
				}
			case "-":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_u - tmp_v)
				} else {
					fmt.Println("no enough arguments to do -")
				}
			case "*":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_u * tmp_v)
				} else {
					fmt.Println("no enough arguments to do *")
				}
			case "/":
				if stack.len > 1 {
					tmp_v := stack.pop()
					tmp_u := stack.pop()
					stack.push(tmp_u / tmp_v)
				} else {
					fmt.Println("no enough arguments to do /")
				}
			default:
				fmt.Println("unrecognized word")
			}
		}
	}
}

func adding_new_keywords(index int, stack *list, arr []string, m map[string][]string) {
	naming := arr[index+1]
	var tmp []string
	for i := index + 2; i < len(arr); i++ {
		if arr[i] == ";" {
			break
		}
		tmp = append(tmp, arr[i])

	}
	if Find(arr, ";") > 0 {
		m[naming] = tmp
		return
	}
	flag := false
	for true {
		if flag == true {
			break
		}
		next := get_array()
		for _, s := range next {
			if s == ";" {
				flag = true
				break
			}
			tmp = append(tmp, s)
		}
	}
	m[naming] = tmp
}

func get_array() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("<Ф ")
	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\r\n")
	arr := split_into_array(str)
	return arr
}
func main() {
	var m map[string][]string
	m = make(map[string][]string)

	stack := initList()
	for true {
		arr := get_array()
		do_switch(arr, stack, m)
	}
}
