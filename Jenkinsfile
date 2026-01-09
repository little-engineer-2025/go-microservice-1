pipeline {
    agent {
        docker {
            image 'docker.io/golang:1.24'
        }
    }
    
    environment {
        DATABASE_HOST = 'postgres'
        DATABASE_PORT = '5432'
        DATABASE_NAME = 'database-db'
        DATABASE_USER = 'database-user'
        DATABASE_PASSWORD = 'database-secret'
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    // Checkout del código y submódulos
                    checkout scm
                }
            }
        }

        stage('Cache Go Modules') {
            steps {
                script {
                    // Usar caché para Go modules
                    // Este paso puede variar según tu configuración de Jenkins
                    sh '''
                        go mod download
                    '''
                }
            }
        }

        stage('Prepare CI') {
            steps {
                script {
                    // Configuración de Git y herramientas
                    sh '''
                        git config --system --add safe.directory $WORKSPACE
                        make tidy
                        make install-go-tools
                    '''
                }
            }
        }

        stage('Run Checks') {
            steps {
                script {
                    // Ejecutar un conjunto de cheques
                    sh '''
                        go version
                        git diff go.mod go.sum tools/go.mod tools/go.sum
                        make generate-api && git diff internal/api/http/
                        make generate-mock && git diff internal/test/mock/
                        make go-fmt && git diff internal/ cmd/
                        make vet
                    '''
                }
            }
        }

        stage('Start Containers') {
            steps {
                script {
                    // Lanza el contenedor de PostgreSQL
                    sh '''
                        make compose-up
                    '''
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    // Preparar archivos de configuración y ejecutar pruebas
                    sh '''
                        cp -vf configs/config.ci.yaml configs/config.yaml
                        make db-migrate-up
                        make test-ci
                    '''
                }
            }
        }

        stage('Stop Containers') {
            steps {
                script {
                    sh '''
                        make compose-down
                    '''
                }
            }
        }

        stage('Generate Code Coverage Report') {
            steps {
                script {
                    // Generar informe de cobertura
                    // La implementación real puede requerir un plugin de Jenkins para la cobertura de código
                    sh '''
                        # Comando para generar cobertura
                    '''
                }
            }
        }

        stage('Add Coverage PR Comment') {
            when {
                expression { return env.BRANCH_NAME ==~ /PR-.*/ } // O ajustar según la lógica de PR
            }
            steps {
                script {
                    // Agregar comentario de cobertura en PR
                    sh '''
                      echo 'TODO: Add comment for cobertura into the PR'
                    '''
                    // Comando para agregar el comentario
                    sh '''
                      echo 'TODO: Push the comment to the PR'
                    '''
                }
            }
        }

        stage('Build Executables') {
            steps {
                script {
                    // Construir ejecutables
                    sh 'make build'
                }
            }
        }
    }

    post {
        always {
            // Limpiar después de la ejecución
            cleanWs()
        }
    }
}

