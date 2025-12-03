pipeline{
    agent any

    stages{
        stage('Checkout'){
            steps{checkout scm}
        }

        stage('Build Go Binary'){
            steps{
                sh 'go mod tidy'
                sh 'go build -o fibservice'
            }
        }

        stage('Docker Build'){
            steps{
                sh 'docker build -t go-fib-service .'
            }
        }

        stage('Run Unit Tests'){
            steps{
                sh 'go test ./...'
            }
        }

        stage('Push to Registry'){
            steps{
                sh 'docker tag go-fib-service aineshsridhar/go-fib-service:latest'
                sh 'docker push aineshsridhar/go-fib-service:latest'
            }
        }

        stage('Deploy'){
            steps{
                sh 'docker-compose down'
                sh 'docker-compose up -d'
            }
        }

        stage('Health Check'){
            steps{
                sh 'curl -f http://localhost:8080/health'
            }
        }
    }
}