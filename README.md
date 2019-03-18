# hello-world-kubernetes

## Pre-reqs
- [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/)
- [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)

## How it works
This project contains:
 - two (really) micro services, hello-micro-service and world-micro-service. Both are written in Go;
 - Dockerfiles for both microservices, both have been pushed to Dockerhub;
 - Kubernetes yaml config files to create a deployment and service;
 - A `setup.sh` script that runs uses the yaml to set up the deployment, pods and a service.

The hello-micro-service runs on port 8080 and has two endpoints. `/hello` simply return "Hello", `/hello-world` returns "Hello world!" by concatenating "Hello" to the result of the endpoint on world-micro-service.
The world-micro-service runs on port 8081 and has one endpoint. `/world` simply returns " world!".

The yaml configures two replica pods containing both microservices, a deployment to manage the pods and a service to load balance to the pods and provide an access point to the application: the IP of the Node on port 30500.

Once you have installed Docker and Minikube, run `setup.sh`, this will set everything up. Note: it may take a few mins for everything to come to life.
You can run `curl $(minikube ip):30500/helloworld` to vertify it all works, the output should be: "Hello world!"
