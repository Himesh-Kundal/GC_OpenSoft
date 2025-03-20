package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"backend/config"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	redisAddr := config.GetEnv("REDIS_ADDR", "localhost:6379")

	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	if err := checkRedisConnection(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func checkRedisConnection() error {
	for i := 0; i < 10; i++ {
		_, err := RedisClient.Ping(ctx).Result()
		if err == nil {
			fmt.Println("Connected to Redis successfully!")
			return nil
		}
		fmt.Println("Redis connection failed. Retrying in 3 seconds...")
		time.Sleep(3 * time.Second)
	}
	return fmt.Errorf("failed to connect to Redis after multiple attempts")
}
