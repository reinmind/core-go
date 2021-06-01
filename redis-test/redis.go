package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := ExampleNewClient()
	ExampleClient(client)
}
func ExampleNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "119.23.225.92:6379",
		Password: "lghcode", // no password set
		DB:       0,         // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
	// Output: PONG <nil>
}
func ExampleClient(client *redis.Client) {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
