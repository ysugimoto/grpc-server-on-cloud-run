Example gRPC server running on Cloud Run.

## Preparation

1. Install `protoc` and `protoc-gen-go`
2. Set up `docker`
3. Set up `gcloud` SDK
4. Set environment `PROJECT_ID` which you want to use GCP project id

## Deployment

Type `make deploy` and it deploys Cloud Run service on full managed service with name of `greeter`.

## Test connection

Install CLI client with nodejs module:

```
npm install -g grpcc
```

And run it:

```
grpcc -p helloworld.proto -a [Cloud Run App host]:443
> client.sayHello({"name": "your name"}, pr)
```

