# gcppubsub
This is the repository for setting gcp pub sub on local machine.

# Usage

## To create subscription and topic
go run main.go --create --projectId=<projectId> --topic=<topic-name> --subscription=<subscription-name>

## To create  topic
go run main.go --create --projectId=<projectId> --topic=<topic-name>

## To delete subscription
go run main.go --delete --projectId=<projectId> --subscription=<subscription-name>

## To delete  topic
go run main.go --delete --projectId=<projectId> --topic=<topic-name>

## To publish message
go run main.go --publish --projectId=<projectId> --topic=<topic-name> --message=<meesage>