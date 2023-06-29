pipeline{
    agent{
        label "any"
    }
    stages{
        stage("Run frontend"){
            steps{
                echo "========executing yarn========"
                nodejs('Node-20.3') {
                    sh 'yarn install'
                }
            }
            post{
                always{
                    echo "========always========"
                }
                success{
                    echo "========A executed successfully========"
                }
                failure{
                    echo "========A execution failed========"
                }
            }
        }
        stage("Run backend"){
            steps{
                echo "========executing gradle========"
                withGradle() {
                    sh './gradlew -v'
                }
            }
            post{
                always{
                    echo "========always========"
                }
                success{
                    echo "========A executed successfully========"
                }
                failure{
                    echo "========A execution failed========"
                }
            }
        }
    }
    post{
        always{
            echo "========always========"
        }
        success{
            echo "========pipeline executed successfully ========"
        }
        failure{
            echo "========pipeline execution failed========"
        }
    }
}