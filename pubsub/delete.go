package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

func DeletePubSubTopic(projectID, topicID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	topic := client.Topic(topicID)
	err = topic.Delete(ctx)
	if err != nil {
		return fmt.Errorf("pubsub topic delete error: %v", err)
	}
	fmt.Println(" topic deleted sucessfully ")
	return nil
}

func DeletePubSubSubscription(projectID, subID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	subs := client.Subscription(subID)
	err = subs.Delete(ctx)
	if err != nil {
		return fmt.Errorf("pubsub subscription delete error: %v", err)
	}
	fmt.Println(" subscription deleted successfully")
	return nil
}
