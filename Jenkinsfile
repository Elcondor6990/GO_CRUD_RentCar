pipeline {
    agent any

    // Usiamo la versione di Go che abbiamo configurato negli strumenti
    tools {
        go 'go-1.26'
    }

    environment {
        // Nome che verrà dato all'immagine finale
        NOME_IMMAGINE = 'mia-api-go'
        GO111MODULE = 'on'
    }

    stages {

        stage('Esecuzione Test') {
            steps {
                echo 'Avvio dei test unitari...'
                sh 'go test -v ./...'
            }
        }

        stage('Build Immagine Docker') {
            steps {
                echo 'Creazione dell immagine Docker tramite il Dockerfile...'
                sh "docker build -t ${NOME_IMMAGINE}:latest ."
            }
        }

        stage('Test Container') {
            steps {
                bat '''
                    docker run --rm ^
                    -p 3000:3000 ^
                    -e MONGO_URI="mongodb://root:pass@mongo:27017/?authSource=admin" ^
                    -e MONGO_DB_NAME="rent_car_crud" ^
                    ${NOME_IMMAGINE}:latest
                '''
            }
        }
    }

    post {
        success {
            echo 'Pipeline completata con successo! L immagine Docker è pronta.'
        }
        failure {
            echo 'La pipeline è fallita. Controlla il Console Output per i dettagli.'
        }
    }
}