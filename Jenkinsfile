// pipeline {
//     agent any

//     environment {
//         // Sesuaikan dengan Docker Registry kamu (misal: dockerhub_username)
//         // Jika hanya main lokal (minikube), bagian push bisa di-skip atau diarahkan ke local registry
//         DOCKER_REGISTRY = "docker.io/dhiyauu"
//         IMAGE_NAME = "Unit_Functional_Test_Tubes2"
//         IMAGE_TAG = "${env.BUILD_ID}"
//     }

//     stages {
//         stage('1. Checkout Repo') {
//             steps {
//                 echo 'Checking out source code...'
//                 checkout scm
//             }
//         }

//         stage('2. Unit Tests') {
//             steps {
//                 echo 'Running Unit Tests (tanpa database)...'
//                 dir('tracking-service') {
//                     // Hanya menjalankan file test yang bukan functional test (jika dipisah dengan build tag)
//                     // Atau secara default go test akan menjalankan unit_test.go yang sudah kita buat
//                     sh 'go test -v ./... -skip Functional'
//                 }
//             }
//         }

//         stage('3. Lint / Vet') {
//             steps {
//                 echo 'Running Go Vet for static analysis...'
//                 dir('tracking-service') {
//                     sh 'go vet ./...'
//                 }
//             }
//         }

//         stage('4. Build Image (Lokal)') {
//             steps {
//                 echo 'Building Docker Image...'
//                 // Build tracking-service image
//                 sh "docker build -t ${IMAGE_NAME}:latest ./tracking-service"
//                 sh "docker tag ${IMAGE_NAME}:latest ${DOCKER_REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
//                 sh "docker tag ${IMAGE_NAME}:latest ${DOCKER_REGISTRY}/${IMAGE_NAME}:latest"
//             }
//         }

//         stage('5. Functional Tests') {
//             steps {
//                 echo 'Running Functional Tests (dengan database)...'
//                 // Menyalakan service dan database menggunakan docker-compose
//                 sh 'docker-compose up -d'
                
//                 // Tunggu sebentar agar database MySQL & service siap menerima koneksi
//                 sleep time: 10, unit: 'SECONDS'
                
//                 dir('tracking-service') {
//                     // Menjalankan spesifik test yang memiliki kata 'Functional' pada namanya
//                     // Note: Test ini dipastikan FAILED sesuai kondisimu saat ini (kode belum selesai)
//                     // Gunakan catchError agar pipeline bisa tetap lanjut / terlihat statusnya
//                     catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
//                         sh 'go test -v -run Functional ./...'
//                     }
//                 }
//             }
//             post {
//                 always {
//                     echo 'Tearing down docker-compose...'
//                     sh 'docker-compose down'
//                 }
//             }
//         }

//         stage('6. Push Image') {
//             steps {
//                 echo 'Pushing image to registry...'
//                 // Pastikan Jenkins sudah login ke registry menggunakan credentials
//                 // sh "echo \$DOCKER_PASSWORD | docker login -u \$DOCKER_USERNAME --password-stdin"
                
//                 sh "docker push ${DOCKER_REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
//                 sh "docker push ${DOCKER_REGISTRY}/${IMAGE_NAME}:latest"
//             }
//         }

//         stage('7. Deploy di Kubernetes') {
//             steps {
//                 echo 'Deploying to Kubernetes...'
//                 // Mengaplikasikan file yaml ke cluster kubernetes
//                 sh 'kubectl apply -f k8s/tracking-deployment.yaml'
//                 sh 'kubectl apply -f k8s/tracking-service.yaml'
                
//                 // Force update image deployment ke tag yang baru saja di-build
//                 sh "kubectl set image deployment/tracking-service tracking-service=${DOCKER_REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
//             }
//         }

//         stage('8. Verify') {
//             steps {
//                 echo 'Verifying deployment...'
//                 // Mengecek status pod dan service untuk memastikan semua berjalan
//                 sh 'kubectl get pods -l app=tracking-service'
//                 sh 'kubectl get svc tracking-service'
                
//                 // Rollout status untuk memastikan deployment selesai tanpa error CrashLoopBackOff
//                 sh 'kubectl rollout status deployment/tracking-service --timeout=60s'
//             }
//         }
//     }

//     post {
//         success {
//             echo 'Pipeline successfully executed!'
//         }
//         failure {
//             echo 'Pipeline failed. Check the logs above.'
//         }
//         unstable {
//             echo 'Pipeline is unstable (Functional test failed as expected).'
//         }
//     }
// }

pipeline {
    agent any

    environment {
        DOCKER_REGISTRY = "docker.io/dhiyauu"
        IMAGE_TAG = "${env.BUILD_ID}"
        KUBECONFIG = "/var/jenkins_home/.kube/config"
    }

    stages {
        stage('1. Checkout Repo') {
            steps {
                echo 'Checking out source code...'
                checkout scm
            }
        }

        stage('2. Unit Tests') {
            steps {
        
                echo 'Running User Service Unit Tests...'
                dir('user-service') {
                    sh 'go test -v ./... -skip Functional'
                }
        
                echo 'Running Order Service Unit Tests...'
                dir('order-service') {
                    sh 'go test -v ./... -skip Functional'
                }
        
                echo 'Running Tracking Service Unit Tests...'
                dir('tracking-service') {
                    sh 'go test -v ./... -skip Functional'
                }
            }
        }

        stage('3. Lint / Vet') {
            steps {

                dir('user-service') {
                    sh 'go vet ./...'
                }

                dir('order-service') {
                    sh 'go vet ./...'
                }

                dir('tracking-service') {
                    sh 'go vet ./...'
                }
            }
        }

        stage('4. Build Images') {
            steps {

                echo 'Building User Service Image...'
                sh 'docker build -t user-service:latest ./user-service'

                echo 'Building Order Service Image...'
                sh 'docker build -t order-service:latest ./order-service'

                echo 'Building Tracking Service Image...'
                sh 'docker build -t tracking-service:latest ./tracking-service'

                // tagging
                sh "docker tag user-service:latest ${DOCKER_REGISTRY}/user-service:${IMAGE_TAG}"
                sh "docker tag order-service:latest ${DOCKER_REGISTRY}/order-service:${IMAGE_TAG}"
                sh "docker tag tracking-service:latest ${DOCKER_REGISTRY}/tracking-service:${IMAGE_TAG}"
            }
        }

       stage('5. Functional Tests') {
            steps {
        
                sh 'docker-compose up -d'
        
                sleep time: 20, unit: 'SECONDS'
        
                catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
        
                    echo 'Running User Functional Tests...'
        
                    dir('user-service') {
                        sh '''
                        docker run --rm \
                          --network tubes2_default \
                          -v $(pwd):/app \
                          -w /app \
                          golang:1.25 \
                          go test -v -run Functional ./...
                        '''
                    }
        
                    echo 'Running Order Functional Tests...'
        
                    dir('order-service') {
                        sh '''
                        docker run --rm \
                          --network tubes2_default \
                          -v $(pwd):/app \
                          -w /app \
                          golang:1.25 \
                          go test -v -run Functional ./...
                        '''
                    }
        
                    echo 'Running Tracking Functional Tests...'
        
                    dir('tracking-service') {
                        sh '''
                        docker run --rm \
                          --network tubes2_default \
                          -v $(pwd):/app \
                          -w /app \
                          golang:1.25 \
                          go test -v -run Functional ./...
                        '''
                    }
                }
            }
        
            post {
                always {
                    sh 'docker-compose down'
                }
            }
        }

        stage('6. Push Images') {
            steps {

                sh "docker push ${DOCKER_REGISTRY}/user-service:${IMAGE_TAG}"
                sh "docker push ${DOCKER_REGISTRY}/order-service:${IMAGE_TAG}"
                sh "docker push ${DOCKER_REGISTRY}/tracking-service:${IMAGE_TAG}"
            }
        }

        stage('7. Deploy Kubernetes') {
            steps {

                sh 'kubectl apply -f k8s/user-deployment.yaml'
                sh 'kubectl apply -f k8s/order-deployment.yaml'
                sh 'kubectl apply -f k8s/tracking-deployment.yaml'

                sh 'kubectl apply -f k8s/user-service.yaml'
                sh 'kubectl apply -f k8s/order-service.yaml'
                sh 'kubectl apply -f k8s/tracking-service.yaml'

                sh "kubectl set image deployment/user-service user-service=${DOCKER_REGISTRY}/user-service:${IMAGE_TAG}"

                sh "kubectl set image deployment/order-service order-service=${DOCKER_REGISTRY}/order-service:${IMAGE_TAG}"

                sh "kubectl set image deployment/tracking-service tracking-service=${DOCKER_REGISTRY}/tracking-service:${IMAGE_TAG}"
            }
        }

        stage('8. Verify') {
            steps {

                sh 'kubectl get pods'
                sh 'kubectl get svc'

                sh 'kubectl rollout status deployment/user-service --timeout=180s'
                sh 'kubectl rollout status deployment/order-service --timeout=180s'
                sh 'kubectl rollout status deployment/tracking-service --timeout=180s'
            }
        }
    }

    post {
        success {
            echo 'Pipeline successfully executed!'
        }
        failure {
            echo 'Pipeline failed. Check the logs above.'
        }
        unstable {
            echo 'Pipeline is unstable (Functional test failed as expected).'
        }
    }
}
