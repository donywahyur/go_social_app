pipeline {
    agent {
        docker {
            image 'golang:1.23'
            args '-v /tmp:/tmp' 
        }
    }
    environment {
        GOPATH = "${WORKSPACE}/go"
        GOBIN = "${GOPATH}/bin"
        PATH = "${GOBIN}:/usr/local/go/bin:${PATH}"
    }
    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/donywahyur/go_social_app.git' 
            }
        }
        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -o app ./cmd'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('Deploy') {
            steps {
                sh './app'  // Run your Go app if needed
            }
        }
    }
}
