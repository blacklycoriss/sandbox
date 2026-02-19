package messaging

import "github.com/rabbitmq/amqp091-go"

var Conn *amqp091.Connection
var Ch *amqp091.Channel

func InitRabbitMQ(url string) {
	Conn, _ = amqp091.Dial(url)
	Ch, _ = Conn.Channel()
	Ch.QueueDeclare("stock-updates", false, false, false, false, nil) // Declare queue
}
