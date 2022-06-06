package main

import (
	"context"
	"fmt"

	"github.com/sy-yoon/krealtors/utils"
)

func main() {
	client := utils.NewRedisClient("192.168.140.130:6379")

	ctx := context.Background()
	pubsub := client.Subscribe(ctx, "region", "realtor")
	defer pubsub.Close()
	//channel := pubsub.Channel()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		switch msg.Channel {
		case "region":
			fmt.Println(msg.Payload)
		case "realtor":
			fmt.Println(msg.Payload)
		}

		fmt.Println(msg.Channel, msg.Payload)
	}
}
