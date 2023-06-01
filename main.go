pipeline {
    agent any
    
    stages {
        stage('Checkout my go project') {
           steps {
               cleanWs()
               git credentialsId: '5d164e7f-1a55-4dd9-a7cf-aa6309cc534b', url:  'https://github.com/Spartan03168/Devops_Go_file.git'
            }
        }
        
        stage('build go') {
            steps {
                sh "go build main.go"
            }
        }
        
        stage('test go') {
            steps {
                sh "go version"
            }
        }
        

        
        stage('target') {
            steps {
            sshagent(['target']) {
                sh "ssh -o StrictHostKeyChecking=no ec2-user@172.31.89.161"
                sh "scp /var/lib/jenkins/workspace/newjob/main ec2-user@172.31.89.161:/home/ec2-user/"
                }
            }
        }
        stage('Deploy go') {
            steps {
                
                sshPublisher(publishers: [sshPublisherDesc(configName: 'mytarget', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: 'nohup ./main &>/dev/null &', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: 'main')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
                
            }
        }
        
    }
}
