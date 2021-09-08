package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"cloud.google.com/go/pubsub"
)

func main() {
	projectID := "benthos-325020"

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	topic := client.Topic("inbound")
	t := time.Now()
	min := 1
	max := 1000
	sorted := rand.Intn(max-min) + min
	message := fmt.Sprintf(`{"id":"%d","names":["celine_%s","dion_%s"]}`, sorted, t.Format("20060102150405"), t.Format("20060102150405"))
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})

	msgID, err := res.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msgID)

}
