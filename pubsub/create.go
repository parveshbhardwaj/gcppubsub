package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

// projects/local-proj-123/topics/sample-project-del-topic
func CreatePubSubTopic(projectID, topicID string) (*pubsub.Topic, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		return nil, fmt.Errorf("CreateTopic: %v", err)
	}
	fmt.Printf("Topic created: %v\n", t)
	return t, nil
}

func CreatePubSubSubscription(projectID, subID string, topic *pubsub.Topic) error {
	// topic of type https://godoc.org/cloud.google.com/go/pubsub#Topic
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub, err := client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 30 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("CreateSubscription: %v", err)
	}
	fmt.Printf("Created subscription: %v\n", sub)
	return nil
}
