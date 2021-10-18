# go-microservices
Learn microservices using Golang

In order to pull all the code for submodules, please execute the following:
```
git submodule init
git submodule update
```

# FOR RUN NPM VUEJS INSTALL
$PWD is the client folder
```
sudo docker run -w /app  -v $PWD/client:/app -p 8080:8080 node /bin/bash -c "npm install -g @vue/cli && vue create -d frontend && npm audit fix"
sudo docker run -w /app  -v $PWD/client:/app -p 8080:8080 -d node /bin/bash -c "cd frontend && yarn serve"
```