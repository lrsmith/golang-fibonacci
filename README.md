# golang-fibonacci

[![Build Status](https://api.travis-ci.org/lrsmith/golang-fibonacci.svg?branch=master)](https://travis-ci.org/lrsmith/go-fibonacci)
[![Go Report Card](https://goreportcard.com/badge/github.com/lrsmith/golang-fibonacci)](https://goreportcard.com/report/github.com/lrsmith/golang-fibonacci)


# Statsd integration

A statsd middleware has been implemented which will send a metric to statsd, for each URI called. The Graphite metric
name is used and defaults to golang-fibonacci.v1.fibseq.<N>  where N is the parameter passed, to indicate the number of Fibonacci numbers to calculate.

Example
`Lens-MacBook-Pro:lrsmith$ curl --insecure -H 'X-Session-Token:00000000' "https://localhost:8443/v1/fibseq?index=8"`
would send a metric to statsd
`golang-fibonacci.v1.fibseq.8:1|c`

# Deployments

 For the initial implementation and deployment, a docker container is created and deployed to an Amazon ECR cluster.

Currently a GNUMakefile is used to control and manage the build and deploy process, with
manual interventions. This should and will be moved to build system, like Jenkins or
Bamboo, where the individual steps would be tasks.

Terraform is used to create and manage AWS infrastructure, such as the AWS ECR
repository and ECS tasks.


## AWS ECR Deployment

1. Make sure the branch you will be deploying from is passing its Travis-CI builds.
2. For production deployments, deploy from 'master' branch and make sure it has been tagged.
3. `make docker-deploy-and-push` This will build a docker container and push to AWS ECR repository.
4. Update terraform with version of Container to deploy and run `terraform apply`
5. Verify

Example
```
Lens-MacBook-Pro:Terraform lrsmith$ curl --insecure -H 'X-Session-Token:00000000' "https://34.228.78.132:8443/v1/fibseq?index=5"
{"httpstatus":200,"sequence":[0,1,1,2,3],"errormsg":""}
```


# To Do
* Add support for configuration management
* Extend `make docker-deploy-and-push` to run integration tests against the docker container before pushing to repository.
* Move Terraform state from local to S3 backed, for easier sharing of state
* Automate updating/creating task-definitions/golang-fibonacci.json by build pipeline for auto deployments to dev,qa environments.
* Use Terraform workspaces to separate state/data between different environments.
* Import ECS cluster information into terraform. Was originally spun up manually
* Make package for generic return message handling.
* Research improving logging, such as source IP, etc.
* Put Load-balancer in front of ECS containers, so can run more than one. Do canary deployments, etc.
* Extend getConfigs to also look for env variables, so can config docke via ENV at runtime, rather then pass in hard-coded configs.
* Add tests for middleware

# ChangeLog

* v0.0.1 - 2018-05-19 : Initial pre-release
