pipeline {
	agent any 

	stages {
//		stage ('Lint') {
//			agent {
//				docker { 
//					image 'golangci/golangci-lint:v1.61.0'
//					reuseNode true
//				}
//			}
//			steps {
//				script {
//					try {
//						sh 'golangci-lint run'
//						echo 'Linting Success!'
//					} catch (err) {
//						echo 'Lint failed'
//						sh 'exit 1'
//					}
//				}
//			}
//		}
		stage ('Build') {
			agent {
				docker { 
					image 'golang:1.23.0'
					reuseNode true
				}
			}
			steps {
				sh 'go build .'
				echo 'Build Success!'
			}
		}
		stage ('Test') {
			steps {
				sh 'docker run --name mysql_${BUILD_NUMBER} -d -e MYSQL_ROOT_PASSWORD=test -e MYSQL_DATABASE=test -e MYSQL_USER=test -e MYSQL_PASSWORD=test -p 3306:3306 mysql'
				sh 'sleep 30'
				sh './main -color=false -dbtable test'
				sh './main -color=false -debug -dbtable test'
				sh './main -json=true -dbtable test'
				sh './main -json=true -debug -dbtable test'
				echo 'Test Success!'
			}
			post {
				always {
					sh 'docker stop --time=1 mysql_${BUILD_NUMBER} '
					sh 'docker rm -vf $(docker ps -aq)'
					echo 'Clean Success!'
					echo ''
					echo 'Type CTRL+C to exit '
				}
			}
		}
	}
}
