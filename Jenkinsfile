ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
   withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
   }
}
pipeline {
    agent any
    environment {
        DOCKER_IMAGE_NAME = "nnungester/go-ape"
    }
    stages {
        stage ('build go app') {
          steps {
            script {
              sh 'export PATH=$PATH:/usr/local/bin/go/bin'
              sh 'go version'
              sh 'go build'
              sh './go-ape'
              sh 'if pgrep go-ape; then pkill go-ape; fi '
              }
            }
          }
        stage('Build Docker Image') {
            steps {
                script {
                    app = docker.build(DOCKER_IMAGE_NAME)
                    app.inside {
                        sh 'echo Hello, World!'
                    }
                }
            }
        }
        stage('Test Docker image with dgoss') {
            steps {
                sh """
                  export GOSS_FILES_STRATEGY=cp && /usr/local/bin/dgoss run -p 8080:8080 nnungester/go-ape
                """
            }
        }
        stage('Push Docker Image') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub-login') {
                        app.push("${env.BUILD_NUMBER}")
                        app.push("latest")
                    }
                }
            }
        }
    }
}
