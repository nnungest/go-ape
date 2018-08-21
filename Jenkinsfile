pipeline {
    agent any
    environment {
        DOCKER_IMAGE_NAME = "nnungester/go-ape"
    }
    stages {
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
                dgoss run -p 8080:8080 nnungester/go-ape
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
