package config

import "os"

func GetDBURL() string       { return os.Getenv("DB_URL") }
func GetRedisAddr() string   { return os.Getenv("REDIS_ADDR") }
func GetRabbitMQURL() string { return os.Getenv("RABBITMQ_URL") }
