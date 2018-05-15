pipeline {
  agent any
  stages {
    stage('') {
      steps {
        sleep(time: 10, unit: 'SECONDS')
        build(job: 'teste', quietPeriod: 1)
        sh 'ls'
      }
    }
  }
}