package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func PublishMessage(projectID, topicID, message string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	topic := client.Topic(topicID)
	var results []*pubsub.PublishResult
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	},
	)
	results = append(results, res)

	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
	return nil
}