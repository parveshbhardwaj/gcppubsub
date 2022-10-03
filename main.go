package main

import (
	"flag"
	"fmt"
	"gcppubsub/pubsub"
	"os"
)

const pubsubHostEnvVar = "PUBSUB_EMULATOR_HOST"

func main() {
	// check for pre-requisite
	if os.Getenv(pubsubHostEnvVar) == "" {
		fmt.Printf("%s environment variable is not set.hence existing the program \n", pubsubHostEnvVar)
		os.Exit(0)
	}
	create := false
	delete := false
	publish := false
	topicID := ""
	projectID := ""
	subscriptionID := ""
	message := ""
	flag.BoolVar(&create, "create", false, "flag for creation")
	flag.BoolVar(&delete, "delete", false, "flag for deletion")
	flag.BoolVar(&publish, "publish", false, "flag for publishing message to pubsub topic")
	flag.StringVar(&topicID, "topic", "", "flag for topic id")
	flag.StringVar(&projectID, "projectId", "", "flag for project Id")
	flag.StringVar(&subscriptionID, "subscription", "", "flag for subscription id")
	flag.StringVar(&message, "message", "", "flag for pub subs message")
	flag.Parse()

	switch {
	case create:
		callCreation(projectID, topicID, subscriptionID)
	case delete:
		callDeletion(projectID, topicID, subscriptionID)
	case publish:
		callPublishMessage(projectID, topicID, message)
	default:
		// write utility help function here
		usage()
	}
}
func usage() {
	fmt.Println(" To use this utility, please choose either of these flags \n --create \n --delete \n --publish")
}
func callCreation(projectID, topicID, subscriptionID string) {
	if len(projectID) == 0 {
		fmt.Println(" please provide project id ")
	}
	if len(topicID) > 0 && len(subscriptionID) == 0 {
		pubsub.CreatePubSubTopic(projectID, topicID)
	} else if len(subscriptionID) > 0 && len(topicID) > 0 {
		topic, err := pubsub.CreatePubSubTopic(projectID, topicID)
		if err != nil {
			fmt.Println(" error in deleting topic ", err)
			return
		}
		pubsub.CreatePubSubSubscription(projectID, subscriptionID, topic)
	}
}

func callDeletion(projectID, topicID, subsID string) {
	if len(projectID) == 0 {
		fmt.Println(" please provide project id ")
		return
	}
	if len(topicID) > 0 {
		pubsub.DeletePubSubTopic(projectID, topicID)
		return
	} else if len(subsID) > 0 {
		pubsub.DeletePubSubSubscription(projectID, subsID)
		return
	}
	fmt.Println(" please provide topic id or subscription id ")
}

func callPublishMessage(projectID, topicID, message string) {
	if len(projectID) == 0 {
		fmt.Println(" please provide project id ")
		return
	}
	if len(topicID) > 0 {
		fmt.Println(" please provide topic id ")
		return
	}
	pubsub.PublishMessage(projectID, topicID, message)
}
