eval $(minikube docker-env)

docker image prune -y

cd ~/Documents/go_projects/auth/
docker build -t auth .
cd ~/Documents/go_projects/people/
docker build -t people .
cd ~/Documents/go_projects/accounts/
docker build -t accounts .

docker images

eval $(minikube docker-env -u)
