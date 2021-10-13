pipeline {
    agent any
    stages{
        stage("Build"){
            steps {
                sh 'cat .gitmodules'
                sh 'git submodule init'
                sh 'git submodule update'
                sh 'docker container prune'
                sh 'docker image prune'
                sh 'docker build -t goauth:latest ./auth/'
                sh 'docker build -t goaccounts:latest ./accounts/'
                sh 'docker build -t gopeople:latest ./people/'

                sh 'docker run --name goauth-container --rm --detach --publish 8081:8080 goauth:latest'
                sh 'docker run --name goauth-container --rm --detach --publish 8081:8080 goaccounts:latest'
                sh 'docker run --name goauth-container --rm --detach --publish 8081:8080 gopeople:latest'
            }            
        }
    }
}