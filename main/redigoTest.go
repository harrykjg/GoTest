package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)
func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "username", "nick")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(c.Do("GET", "username"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Got username %v \n", username)
	}



	_, err = c.Do("SET", "password", "123456", "EX", 5)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	time.Sleep(6 * time.Second)
	password, err := redis.String(c.Do("GET", "password"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Got password %v \n", password)
	}



	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")
	c.Flush()
	//c.Receive() // reply from SET
	v, err := c.Receive() // reply from GET
	fmt.Println(v)
}
