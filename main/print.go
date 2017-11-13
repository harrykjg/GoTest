package main

import (
	"fmt"
	"os/user"
)

func main(){

	k :=10
	if k>9{
		for i:=1;i<10;i++{
			if i==5{
				break
			}
			fmt.Println(i);


		}
	}
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println( usr.HomeDir )


}
