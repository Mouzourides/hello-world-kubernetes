#!/bin/bash

kubectl create -f deployment.yml
kubectl create -f service.yml

sleep 20

curl $(minikube ip):30500/helloworld
