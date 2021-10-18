#!/bin/bash

#eval $(minikube docker-env)

docker container stop accounts
docker container stop auths
docker container stop people
docker container stop mysql

docker network prune -f
docker container prune -f 
docker image prune -f

docker build -t auth ./auth
docker build -t accounts ./accounts
docker build -t people ./people

docker run --detach --name accounts --env-file env/variables -p 8081:8080  accounts
docker run --detach --name auths --env-file env/variables -p 8082:8080 auth
docker run --detach --name people --env-file env/variables  -p 8083:8080 people
docker run --detach --name mysql -p 3306:3306 --env-file env/variables mysql:latest

docker images
docker ps --all

python3 -m http.server

#eval $(minikube docker-env -u)
