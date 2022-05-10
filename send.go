package main

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"context"
	"fmt"
	"encoding/json"
)

type MessageInfo struct {
	Body   string
	Result *primitive.SendResult
}

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

	msgBody := "Hello RocketMQ Go Client!"
	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(msgBody),
	}
	msg.WithDelayTimeLevel(0)
	// msg.WithTag("")
	// msg.WithKeys([]string{""})

	result, err := rocketmqProducer.SendSync(context.Background(), msg)
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(MessageInfo{
		Body:   msgBody,
		Result: result,
	})
	fmt.Println(string(b))
}
