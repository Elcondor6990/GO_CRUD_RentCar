pipeline {
    agent any 

    stages {

        stage('Build Go') {
            agent { label 'jenkins-agent-golang' } 
            steps {
                echo "Building Go application..."
                sh 'go version'
                sh 'go mod tidy'
                sh 'go build -o app'
            }
        }

        stage('Test Go') {
            agent { label 'jenkins-agent-golang' }
            steps {
                echo "Running Go tests..."
                sh 'go test ./...'
            }
        }

        stage('Build Docker Image') {
            agent { label 'jenkins-agent-docker' } 
            steps {
                echo "Building Docker image..."
                sh 'docker build -t my-go-app:latest .'
            }
        }

        stage('Run Docker Container') {
            agent { label 'jenkins-agent-docker' }
            steps {
                echo "Running Docker container..."
                sh '''
                    docker rm -f my-go-app || true
                    docker run --rm -d --name my-go-app -p 8081:8080 my-go-app:latest
                '''
            }
        }

        stage('Kubectl Test') {
            agent { label 'kubectl-agent' }
            steps {
                withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
                    sh 'ls -l $KUBECONFIG'
                    sh 'kubectl cluster-info'
                    sh 'kubectl get nodes'
                }
            }
        }
    }

    post {
        always {
            echo 'Pipeline finished'
        }
    }
}