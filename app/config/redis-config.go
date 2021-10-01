package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func SetupRedisConnection() *redis.Client {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	dns := os.Getenv("REDIS_URI")
	pass := os.Getenv("REDIS_PASS")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     dns,
		Password: pass, // no password set
		DB:       db,  // use default DB
	})

	return rdb
}