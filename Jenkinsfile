pipeline {
    agent any
    stages {
        stage('Environment') {
            steps {
                echo 'Populating environment variables..'
                sh """
                echo "APP_ADDRESS=${env.APP_ADDRESS}" > .env
                echo "APP_PORT=${env.APP_PORT}" >> .env
                echo "DB_DSN=${env.DB_DSN}" >> .env
                echo "DB_USER=${env.DB_USER}" >> .env
                echo "DB_PASSWORD=${env.DB_PASSWORD}" >> .env
                echo "DB_PORT=${env.DB_PORT}" >> .env
                echo "DB_NAME=${env.DB_NAME}" >> .env
                echo "DB_MAX_OPEN=${env.DB_MAX_OPEN}" >> .env
                echo "DB_MAX_IDLE=${env.DB_MAX_IDLE}" >> .env
                echo "DB_MAX_LIFE=${env.DB_MAX_LIFE}" >> .env
                echo "JWT_SECRET=${env.JWT_SECRET}" >> .env
                echo "JWT_EXPIRE=${env.JWT_EXPIRE}" >> .env
                """
                echo 'Environment variables populated.'
            }
        }
        stage('Build') {
            steps {
                echo 'Building docker image..'
                sh 'docker build --platform=linux/amd64 -t mfathoor/posyandu-api:latest .'
                echo 'Docker image built.'
            }
        }
        stage('Push') {
            steps {
                echo 'Pushing docker image..'
                withCredentials([usernamePassword(credentialsId: 'fathoor-docker-hub', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                    sh 'docker login -u $USERNAME -p $PASSWORD'
                }
                sh 'docker push mfathoor/posyandu-api:latest'
                echo 'Docker image pushed.'
            }
        }
        stage('Deploy'){
            stages {
                stage('Down') {
                    steps {
                        echo 'Stopping and removing existing container..'
                        sh 'docker compose down'
                        echo 'Existing container stopped and removed.'
                    }
                }
                stage('Clean') {
                    steps {
                        echo 'Removing existing image..'
                        sh 'docker image rm mfathoor/posyandu-api:latest'
                        echo 'Existing image removed.'
                    }
                }
                stage('Up') {
                    steps {
                        echo 'Starting new container..'
                        sh 'docker compose up -d'
                        echo 'New container started.'
                    }
                }
            }
        }
    }
    post {
        success {
            echo 'App is deployed successfully.'
        }
        failure {
            echo 'App deployment failed.'
        }
        cleanup {
            echo 'Cleaning up..'
            deleteDir()
            echo 'Cleaned up.'
        }
    }
}
