pipeline{
    agent any

    environment{
        IMAGE_NAME = "aineshsridhar/go-fib-service:latest"
    }

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
        
        stage('Run Unit Tests'){
            steps{
                sh 'go test ./...'
            }
        }

        stage('Docker Build'){
            steps{
                sh 'docker build -t ${IMAGE_NAME} .'
            }
        }

        stage('Push to Registry'){
            steps{
                sh 'docker push ${IMAGE_NAME}'
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl apply -f deployment.yaml'
                sh 'kubectl apply -f service.yaml'
                sh 'kubectl apply -f hpa.yaml'
            }
        }

        stage('Health Check'){
            steps{
                sh 'kubectl rollout status deployment/go-fib-deploy'
            }
        }
    }
}