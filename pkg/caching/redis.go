package caching

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type Client struct {
	*redis.Client
}

var RedisClient Client

func Initialize() Client {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"), 				 // Replace with your Redis server address
		Password: viper.GetString("password.address"),               // Replace with your Redis server password
		DB:       viper.GetInt("password.db"),                       // Replace with your Redis database number
	})

	// Ping the Redis server to check the connection
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Redis:", pong)
	RedisClient = Client{Client: client}
	return RedisClient
}

