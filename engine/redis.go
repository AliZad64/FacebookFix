package engine

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
)

func InitRedis() {
	ctx := context.Background()
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	RDB = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       redisDB,
	})
	err := RDB.Ping(ctx).Err()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = RDB.ConfigSet(ctx, "notify-keyspace-events", "KEx").Err()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("connected to redis at %s", redisHost)
}
