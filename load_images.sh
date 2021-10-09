eval $(minikube docker-env)

docker image prune -y


docker build -t auth ./auth

docker build -t people ./people

docker build -t accounts ./accounts

docker images

eval $(minikube docker-env -u)
