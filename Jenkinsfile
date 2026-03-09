pipeline {
    agent none  // nessun agent globale

    stages {

        stage('Build Go') {
            agent { label 'golang' } // usa l'agent Go
            steps {
                echo "Building Go application..."
                sh 'go version'
                sh 'go mod download'
                sh 'go build -o app'
            }
        }

        stage('Test Go') {
            agent { label 'golang' }
            steps {
                echo "Running Go tests..."
                sh 'go test ./...'
            }
        }

        stage('Build Docker Image') {
            agent { label 'docker' } // usa l'agent Docker
            steps {
                echo "Building Docker image..."
                sh 'docker build -t my-go-app:latest .'
            }
        }

        stage('Run Docker Container') {
            agent { label 'docker' }
            steps {
                echo "Running Docker container..."
                sh '''
                    # Rimuove eventuali container esistenti
                    docker rm -f my-go-app || true
                    docker run --rm -d --name my-go-app -p 8081:8080 my-go-app:latest
                '''
            }
        }
    }

    post {
        always {
            echo 'Pipeline finished'
        }
    }
}