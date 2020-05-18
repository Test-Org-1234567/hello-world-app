#!/usr/bin/env groovy

def label = "k8sagent-e2el"
def home = "/home/jenkins"
def workspace = "${home}/workspace/build-jenkins-operator"
def workdir = "${workspace}/src/github.com/jenkinsci/kubernetes-operator/"

podTemplate(label: label, yaml: """
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: docker-dind
    image: docker:18-dind
    securityContext:
      privileged: true
    env:
    - name: HTTPS_PROXY
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: https_proxy
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: HTTP_PROXY
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: http_proxy
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: NO_PROXY
      value: localhost,127.0.0.1
    - name: no_proxy
      value: localhost,127.0.0.1
  - name: docker
    image: docker:18.06.3-ce
    command: ['cat']
    tty: true
    env:
    - name: DOCKER_HOST
      value: tcp://localhost:2375 
    - name: HTTPS_PROXY
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: https_proxy
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: HTTP_PROXY
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: http_proxy
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: NO_PROXY
      value: localhost,127.0.0.1
    - name: no_proxy
      value: localhost,127.0.0.1
  - name: golang
    image: golang:1.13
    command: ['cat']
    tty: true
    env:
    - name: HTTPS_PROXY
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: https_proxy
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: HTTP_PROXY
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: http_proxy
      value: http://proxy.ftc5.hpelabs.net:8080
    - name: NO_PROXY
      value: localhost,127.0.0.1
    - name: no_proxy
      value: localhost,127.0.0.1
"""
  ) {
    node(label) {

        if (env.TAG_NAME != null) {
        stage('Checkout Workspace') {
          checkout([$class: 'GitSCM', branches: [[name: "refs/tags/$TAG_NAME"]], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'SubmoduleOption', disableSubmodules: false, parentCredentials: false, recursiveSubmodules: false, reference: '', trackingSubmodules: false]], submoduleCfg: [], userRemoteConfigs: [[url: "https://github.com/Test-Org-1234567/hello-world-app"]]])
        }
        } else {
        stage('Checkout Workspace') {
            git url: 'https://github.com/Test-Org-1234567/hello-world-app', branch: env.BRANCH_NAME
        }
        }
        stage('Dependencies') {
          container('golang') {
            sh 'go mod download'
          }            
        }
        stage('Tests') {
             parallel (
               lint: {
                     container('golang') {
			 sh 'apt-get update'
		         sh 'apt-get install -y curl'
                         sh '[ -f "$(which golangci-lint)" ] || (curl -fsSL -o /tmp/linter.tar.gz https://github.com/golangci/golangci-lint/releases/download/v1.17.1/golangci-lint-1.17.1-linux-amd64.tar.gz && mkdir -p /tmp/linter && tar xzf /tmp/linter.tar.gz --strip-components=1 -C /tmp/linter && mv /tmp/linter/golangci-lint /usr/local/bin/)'
                         sh 'echo Running golangci-lint && golangci-lint run --config .golangci.yaml'
                     }            
              },
               unit_test: {
                     container('golang') {
                         sh 'go test ./...'
                     }            
              })
        }
        stage('Build') {
           container('golang') {
             sh 'go build -o main cmd/hello-world/main.go'
           }            
        }
        stage('Docker build') {
            container('docker') {
                sh 'docker build -t hardy047/hello-world-app:"${BRANCH_NAME}" .'
            }
        }
        stage('Functional Test') {
            container('docker') {
                 sh '''
                   #apk add curl
                   #curl -Lo /usr/local/bin/kind https://github.com/kubernetes-sigs/kind/releases/download/v0.5.1/kind-$(uname)-amd64
                   #chmod +x /usr/local/bin/kind
                   #curl -Lo /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v1.15.3/bin/linux/amd64/kubectl
                   #chmod +x /usr/local/bin/kubectl
                   #kind create cluster --wait 5m
                   #export KUBECONFIG=$(kind get kubeconfig-path)
                   #kind get clusters
                   #kubectl get pod --all-namespaces
                 '''
            }
        }
        if (env.TAG_NAME != null) {
        stage('Publish') {
            container('docker') {
                withDockerRegistry([ credentialsId: "dockerhub-login", url: "" ]) {
	          sh 'docker push hardy047/hello-world-app:"${BRANCH_NAME}"'
                }
            }
          }
        }
    }
}
