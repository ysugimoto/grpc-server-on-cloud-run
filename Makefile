.PHONY: build deploy

build:
	protoc --go_out=plugins=grpc:./spec helloworld.proto

deploy: build
	gcloud builds submit --tag asia.gcr.io/${PROJECT_ID}/greeter
	gcloud beta run deploy greeter --image asia.gcr.io/${PROJECT_ID}/greeter --platform managed --region asia-northeast1


