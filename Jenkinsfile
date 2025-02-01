pipeline {
    agent {
        docker {
            image 'golang:1.20'  // Use a valid Go version
            args '-v /tmp:/tmp'  // Optional: Remove if not needed
        }
    }
    environment {
        GOPATH = "${WORKSPACE}/go"
        GOBIN = "${GOPATH}/bin"
        PATH = "${GOBIN}:/usr/local/go/bin:${PATH}"
        DB_HOST= credentials('DB_HOST')
        DB_USER= credentials('DB_USER')
        DB_PASS= credentials('DB_PASS')
        DB_NAME= credentials('DB_NAME')
        DB_PORT= credentials('DB_PORT')
    }
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/donywahyur/go_social_app.git'  // Specify branch
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
                sh './app' 
            }
        }
    }
    post {
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}