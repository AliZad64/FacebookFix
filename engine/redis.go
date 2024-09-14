package engine

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
	log.Printf("connected to redis at %s", redisHost)
}

func SetRedisContent(ctx *gin.Context, key, value string) error {
	err := RDB.Set(ctx, key, value, 7*time.Hour).Err()
	return err
}

func GetRedisContent(ctx *gin.Context, key string) (string, error) {
	val, err := RDB.Get(ctx, key).Result()
	return val, err
}
