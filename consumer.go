package main

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"context"
	"fmt"
	"time"
)

func main() {
	topic := "testTopic"

	rocketmqPushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"0.0.0.0:9876"})),
	)
	if err != nil {
		panic(err)
	}

	err = rocketmqPushConsumer.Subscribe(
		topic,
		consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				fmt.Println("===")
				fmt.Println(fmt.Sprintf("subscribe callback: %v", string(msgs[i].Body)))
				fmt.Println("===")
			}

			return consumer.ConsumeSuccess, nil
		},
	)
	if err != nil {
		panic(err)
	}

	// Note: start after subscribe
	err = rocketmqPushConsumer.Start()
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Hour)
	err = rocketmqPushConsumer.Shutdown()
	if err != nil {
		panic(err)
	}
}
