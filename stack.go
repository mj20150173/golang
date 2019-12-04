package main

import (
	"fmt"


)


type Node struct{
	value int
}

func (n *Node) String() string{
	return fmt.Sprint(n.value)
}

func NewStack() *Stack{
	return &Stack{}
}
type Stack struct{
	nodes []*Node
	count int
}

func (s *Stack) Push(n *Node){
	s.nodes=append(s.nodes[:s.count],n)
	s.count++
}
func (s *Stack) Pop() *Node {
	if s.count==0{
		fmt.Println("데이터가 존재하지 않음!")
		return nil
	}
	s.count--
	return s.nodes[s.count]
}


func main() {

	s:=NewStack()

	s.Push(&Node{1})
	s.Push(&Node{2})
	s.Push(&Node{3})
	s.Push(&Node{4})
	s.Push(&Node{5})
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
