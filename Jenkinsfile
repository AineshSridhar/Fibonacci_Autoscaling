pipeline{
    agent any

    environment{
        IMAGE_NAME = "aineshsridhar/go-fib-service:latest"
        KUBECONFIG = "/var/lib/jenkins/.kube/config"
    }

    stages{
        stage('Checkout'){
            steps{checkout scm}
        }

        stage('Build Go Binary'){
            steps{
                sh 'go build -mod=vendor -o fibservice'
            }
        }
        
        stage('Run Unit Tests'){
            steps{
                sh 'go test ./...'
            }
        }

        stage('Docker Build'){
            steps{
                sh '''
                    eval $(minikube -p minikube docker-env)
                    docker build -t ${IMAGE_NAME} .'
                '''
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl apply --validate=false -f deployment.yaml'
                sh 'kubectl apply --validate=false -f service.yaml'
                sh 'kubectl apply --validate=false -f hpa.yaml'
            }
        }

        stage('Health Check'){
            steps{
                sh 'kubectl rollout status deployment/go-fib-deploy'
            }
        }
    }
}