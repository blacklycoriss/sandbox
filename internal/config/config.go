package config

import "os"

func GetDBURL() string       { return os.Getenv("DB_URL") } // postgres://...
func GetRedisAddr() string   { return "redis:6379" }
func GetRabbitMQURL() string { return "amqp://guest:guest@rabbitmq:5672/" }
