# cdk-demo

This little demo was written to test out CDK using GoLang.

I wanted to validate the following:
* service that could be run locally
* deployable to lambda, servicing requests from API Gateway
* using CDK to deploy everything in one step

This demo proves the above out nicely.

I've used the [Gorilla mux router](https://github.com/akrylysov/algnhsa) and [algnhsa](https://github.com/akrylysov/algnhsa) AWS Lambda/http compatability layer.

## Getting Started

* Configure your AWS Credentials
* Install AWS CDK 
  * `npm install -g aws-cdk`
* Bootstrap CDK:
  * `cdk bootstrap aws://<account-id>/eu-west-1`
  * This sets up an S3 bucket and various stuff to get you started
* Build the app!
  * `go mod download && cd api && go build`
* Run locally to test if desired
  * `go run api/main.go`
  * `curl -XPOST --data '{"email": "test"}' localhost:8000`
* Deploy the CDK stack
  * I needed to setup my profile and region first `export AWS_PROFILE=xxx && export AWS_REGION=eu-west-1`
  * `cdk deploy`
* Test it out!
  * `curl -XPOST --data '{"email": "test"}' <URL CDK gives you>/welcome`