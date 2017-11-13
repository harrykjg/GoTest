package main

import "fmt"


type Origin struct{

	name string
}
type F interface{
	F1 (string) int
	F2 (string) string
}

func (o *Origin ) F1 (s string) int{
	return 1
}

func (o *Origin ) F2 (s string) string{
	return "sb"
}

type Mock struct {
	mockf1 func() int
}
func (m *Mock ) F1 (s string) int{
	//if(m.mockf1 !=nil){
	//	return m.mockf1()
	//}
	//return 0
	oo:=&Origin{}
	return oo.F1("s")
}

func (m *Mock ) F2 (s string) string{
	return "DSB"
}

func real (fp F) string{
	s1:=fp.F1("ss")
	s2:=fp.F2("ss")
fmt.Println(s1,s2)
	return "xx"
}
func main(){
	var o=&Origin{}
	var m=&Mock{
		//mockf1: func() int {
		//	return 2
		//},
	}
	fmt.Println(real(o))
	fmt.Println(real(m))
}
