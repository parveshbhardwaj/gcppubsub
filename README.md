# gcppubsub
This is the repository for setting gcp pub sub on local machine.

# Usage

## Prerequisite
Install Pubsub on your local using below link

https://cloud.google.com/pubsub/docs/emulator

set up the PUBSUB_EMULATOR_HOST environment variable

## To create subscription and topic
go run main.go --create --projectId=projectId --topic=topic-name --subscription=subscription-name

## To create  topic
go run main.go --create --projectId=projectId --topic=topic-name

## To delete subscription
go run main.go --delete --projectId=projectId --subscription=subscription-name

## To delete  topic
go run main.go --delete --projectId=projectId --topic=topic-name

## To publish message
go run main.go --publish --projectId=projectId --topic=topic-name --message=message