// Go supports _methods_ defined on struct types.

package main

import "fmt"

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *rect) area() int {               //*rect or rect is the same
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {                  //*rect or rect is the same
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	r2:=rect{width:20,height:5}
        fmt.Println(r2)
	// Go automatically handles conversion between values    //参考这个网站http://studygolang.com/articles/4059,说把方法赋给指针和值的区别
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	rp := &r     //&r and r is the same
	fmt.Println(rp)
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
