package main

import (
	"fmt"
	"time"

	"gopkg.in/go-redis/cache.v1"
	"gopkg.in/redis.v3"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type Object struct {
	Str string

}
func main(){






	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":uaa_url1",
		},

		DialTimeout:  3 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	})

	codec := &cache.Codec{
		Ring: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	key := "mykey"
	obj := &Object{
		Str: "mystring",
	}

	codec.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: time.Hour,
	})

	var wanted Object
	if err := codec.Get(key, &wanted); err == nil {
		fmt.Println(wanted)
	}


	if err := codec.Get(key, &wanted); err == nil{
             fmt.Println(wanted)
	}

}
