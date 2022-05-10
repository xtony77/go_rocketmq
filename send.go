package main

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"context"
	"fmt"
)

func main() {
	topic := "testTopic"

	rocketmqProducer, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"0.0.0.0:9876"})),
		producer.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}

	err = rocketmqProducer.Start()
	if err != nil {
		panic(err)
	}

	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte("Hello RocketMQ Go Client!"),
	}
	msg.WithDelayTimeLevel(0)
	result, err := rocketmqProducer.SendSync(context.Background(), msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
