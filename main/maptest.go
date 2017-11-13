package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type DeploymentStatus struct {
	ID        string    `json:"Id"`
	Org       string    `json:",omitempty"`

	Request   PushRequest

}
type PushRequest struct{
	A string
	B string `json:"B"`
	E Envelop
}

type Envelop struct{
	X string
}

func main() {
        s:= DeploymentStatus{}
	s.ID="id"
	s.Org="org"
	e:=Envelop{}
	e.X="xxx"
	r:=PushRequest{"a","b",e}
	s.Request=r
	fmt.Println(s)
	ma, err := json.Marshal(s)    //首先得理解这个marshal和unmarshal是干嘛,marshal先把这个object(或者是
	                              // b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)) encode成json编码的byte数组,
	if err != nil {               //有了这个byte数组之后就可以把它转成map或者转回原来的那个对象,或者直接把这个byte数组写出成response
		fmt.Println("fali marshal")
	}
	fmt.Println(ma)
	var dat map[string]interface{}
	json.Unmarshal(ma,&dat)
	fmt.Println(dat)
	fmt.Println(dat["Id"])//这个转成map之后就是取那个`json:"Id"`作为key
	var s2 DeploymentStatus
	json.Unmarshal(ma,&s2)
	fmt.Println(s2.Request)
	//b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000,"abc"{"k":"kk"}}`)

	//fmt.Println(ma["Request"]["Envelop"])
	//var dat map[string]interface{}
	//if err := json.Unmarshal(ma, &dat); err != nil {
	//
	//}
	//fmt.Println(dat["Request"]["Envelop"])
	//dat["Request"]["Envelop"]=""
	//fmt.Println(dat["Request"]["Envelop"])

fmt.Println("测试如何把map输出到file")
	//var m2 map[string]interface{} =make(map[string]interface{})
	//m2["x"]="b"
	//ma2, err2:= json.Marshal(m2)
	//fmt.Println(err2)
	//err3 := ioutil.WriteFile("dat4.yml", ma2, 0777)
	////
	//fmt.Println(err3)

	var m2 map[string]interface{} =make(map[string]interface{})
	m2["x"]="| b"
	ma2, err2:= yaml.Marshal(m2)
	fmt.Println(err2)

	err3 := ioutil.WriteFile("dat4.yml", ma2, 0777)
	//
	fmt.Println(err3)
}
