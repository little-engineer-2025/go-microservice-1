// https://www.jenkins.io/doc/book/pipeline/syntax/#declarative-pipeline
pipeline {
    agent any
    
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

        stage('Generate caches') {
            agent { docker { image 'docker.io/golang:1.24' } }
            steps {
                script {
                    // Generate cache
                    cache(maxCacheSize: 500, caches: [
                        arbitraryFileCache(path: '/github/home/.cache/go-build', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: '/go/pkg/mod', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: './tools/bin', cacheValidityDecidingFile: 'tools/go.sum'),
                        arbitraryFileCache(path: './.venv', cacheValidityDecidingFile: 'requirements-dev.txt')
                    ]) {
                        make tidy
                        make install-go-tools
                    }
                }
            }
        }

        stage('Run Checks') {
            agent { docker { image 'docker.io/golang:1.24' } }
            steps {
                script {
                    cache(maxCacheSize: 500, caches: [
                        arbitraryFileCache(path: '/github/home/.cache/go-build', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: '/go/pkg/mod', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: './tools/bin', cacheValidityDecidingFile: 'tools/go.sum'),
                        arbitraryFileCache(path: './.venv', cacheValidityDecidingFile: 'requirements-dev.txt')
                    ], skipSave: true) {
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
        }

        stage('Start Containers') {
            steps {
                script {
                    sh '''
                        make compose-up
                    '''
                }
            }
        }

        stage('Run Tests') {
            agent { docker { image 'docker.io/golang:1.24' } }
            steps {
                script {
                    cache(maxCacheSize: 500, caches: [
                        arbitraryFileCache(path: '/github/home/.cache/go-build', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: '/go/pkg/mod', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: './tools/bin', cacheValidityDecidingFile: 'tools/go.sum'),
                        arbitraryFileCache(path: './.venv', cacheValidityDecidingFile: 'requirements-dev.txt')
                    ], skipSave: true) {
                        sh '''
                            cp -vf configs/config.ci.yaml configs/config.yaml
                            make db-migrate-up
                            make test-ci
                        '''
                    }
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
            agent { docker { image 'docker.io/golang:1.24' } }
            steps {
                script {
                    cache(maxCacheSize: 500, caches: [
                        arbitraryFileCache(path: '/github/home/.cache/go-build', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: '/go/pkg/mod', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: './tools/bin', cacheValidityDecidingFile: 'tools/go.sum'),
                        arbitraryFileCache(path: './.venv', cacheValidityDecidingFile: 'requirements-dev.txt')
                    ], skipSave: true) {
                        // Generar informe de cobertura
                        // La implementación real puede requerir un plugin de Jenkins para la cobertura de código
                        sh '''
                            # TODO Comando para generar cobertura
                        '''
                    }
                }
            }
        }

        stage('Add Coverage PR Comment') {
            when {
                expression { return env.BRANCH_NAME ==~ /PR-.*/ } // O ajustar según la lógica de PR
            }
            agent { docker { image 'docker.io/golang:1.24' } }
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
            agent { docker { image 'docker.io/golang:1.24' } }
            steps {
                script {
                    cache(maxCacheSize: 500, caches: [
                        arbitraryFileCache(path: '/github/home/.cache/go-build', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: '/go/pkg/mod', cacheValidityDecidingFile: 'go.sum'),
                        arbitraryFileCache(path: './tools/bin', cacheValidityDecidingFile: 'tools/go.sum'),
                        arbitraryFileCache(path: './.venv', cacheValidityDecidingFile: 'requirements-dev.txt')
                    ], skipSave: true) {
                        // Construir ejecutables
                        sh 'make build'
                    }
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

