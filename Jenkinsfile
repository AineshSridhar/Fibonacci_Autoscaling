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
                sh 'go build -mod=vendor -o fib-service'
            }
        }
        
        stage('Run Unit Tests'){
            steps{
                sh 'go test ./...'
            }
        }

        stage('Docker Build & Push'){
            steps{
                withCredentials([usernamePassword(credentialsId: 'docker-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]){
                    sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                    sh 'docker build -t ${IMAGE_NAME} .'
                    sh 'docker push ${IMAGE_NAME}'
                    sh 'docker system prune -f'
                }
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