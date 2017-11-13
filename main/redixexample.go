package main

import (
	"fmt"
	"gopkg.in/redis.v3"
)

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	err2 := client.Set("key", "value", 0).Err()
	if err2 != nil {
		panic(err2)
	}

	val, err2 := client.Get("key").Result()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("key", val)

	val2, err2 := client.Get("key2").Result()
	if err2 == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err2 != nil {
		panic(err2)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists
}

func main(){
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	err2 := client.Set("key", "value", 0).Err()
	if err2 != nil {
		panic(err2)
	}

	val, err2 := client.Get("key").Result()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("key", val)

	val2, err2 := client.Get("key2").Result()
	if err2 == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err2 != nil {
		panic(err2)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists

}
