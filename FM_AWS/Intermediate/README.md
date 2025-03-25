# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests

## Deployment steps

1. `cd lambda`       navigate to the lambda directory
2. `make build`      create the zip bootstrap
3. `cd ..`           navigate back to the Intermediate directory
4. `cdk deploy`      deploy your stack to AWS

## API Endpoints

#### > Make sure to update the URL, username, password and authorization bearer fields as needed.

### Register

To register a new user, use the following `curl` command:

```sh
curl -X POST https://61tgx5b10h.execute-api.us-east-1.amazonaws.com/prod/register -H "Content-Type: application/json" -d '{"username":"Kalki","password":"Shivay123"}'
```

### Login
To log in, use the following curl command:

```bash
curl -X POST https://61tgx5b10h.execute-api.us-east-1.amazonaws.com/prod/login -H "Content-Type: application/json" -d '{"username":"Kalki","password":"Shivay123"}'
```

### Access Protected Resource
To access a protected resource, use the following curl command:

```bash
curl -X GET https://61tgx5b10h.execute-api.us-east-1.amazonaws.com/prod/protected -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoxNzQyOTE5ODEyLCJ1c2VyIjoiS2Fsa2kifQ.7MeCZhHpdpZtPlW02-mcp101CGKnJhY_soUm67KJ8II"
```