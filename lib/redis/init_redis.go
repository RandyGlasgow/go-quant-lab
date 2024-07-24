package redis

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

type RedisRepository struct {
	Client *redis.Client
}

var Redis RedisRepository

func InitRedis() {
	godotenv.Load(".env")

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))

	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(opt)

	_, err = client.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal(err)
	}

	Redis = RedisRepository{Client: client}
}
