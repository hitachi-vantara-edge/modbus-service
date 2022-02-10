library identifier: 'shared-jenkinsfile-pipelines@master', retriever: modernSCM(
  [$class: 'GitSCMSource',
   remote: 'https://github.com/lumada-edge/devops_shared-jenkinsfile-pipelines.git',
   credentialsId: '6e5c3081-96ac-4a13-9c02-5af7ea7bcabd'])

// --------------- CHANGE THESE TO FIT YOUR PROJECT ---------------
// The name of your repo
def project_name = "modbus-service"
// The path in GitHub
def project_dir = "lumada-edge/"
// Domain (if necessary)
def project_domain = "github.com/"
// Git SSH
def git_url = "https://github.com/hitachi-vantara-edge/modbus-service.git"
def artifactory_url="lumadaedge-docker-dev-sc.repo.sc.eng.hitachivantara.com"
// Name you wish to give your binary

// Image name (should begin with hiota)
def image_name = "modbus-service"
def image_tag = ""
def modbus_lib_version = "3.1.4"
// Source code directory

// Dockerfile name for building your image
// ----------------------------------------------------------------

def build_branch = "$build_branch"
def docker_image = null
def commit_author = null
def commit_message = null
def commit_sha = null
project_path = "/src/$project_domain$project_dir$project_name"
def failure_step = ""

node {
    try {

        // Needed for trigger release builds or general pipeline builds
        if (image_tag == "") {
            image_tag = currentBuild.id
        }

        checkout scm

        // Build docker image to build binary in
        docker.withRegistry("http://${artifactory_url}", 'artifactory') {
            stage('Prepare Docker Image') {

                failure_step = "Prepare Docker Image"
                docker_image_for_arm64_build = docker.image("${artifactory_url}/repository/pandora/build_images/multarch-builder:experiment-01")
            }
        }

        // Work in a docker image that we will run in
        docker_image_for_arm64_build.inside('--privileged -v /var/run/docker.sock:/var/run/docker.sock -v /etc/docker:/etc/docker -u root -v $WORKSPACE:/output --add-host=lumada-edge.hiota.pikachu.lumadaedge:10.76.11.131 --add-host=pikachu.lumadaedge:10.76.11.131 --add-host=lumada-edge.hiota.torchic.lumadaedge:10.76.11.171 --add-host=torchic.lumadaedge:10.76.11.171') {
            // Get the Go Stuff
            stage('Checkout Code') {

                failure_step = "Checkout Code"

                sh "mkdir -p $project_path"

                checkout([$class: 'GitSCM', branches: [[name: build_branch]], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: "$project_name"]], submoduleCfg: [], userRemoteConfigs: [[credentialsId: '6e5c3081-96ac-4a13-9c02-5af7ea7bcabd', url: "$git_url"]]])

                commit_author = CommitAuthor("$project_name")
                commit_message = CommitMessage("$project_name")
                commit_sha = CommitHash("$project_name")

                sh "mv $project_name/* $project_path"
                sh "rm -Rf $project_name"
            }

            stage('Building AMD64 and ARM64 Images') {
                failure_step = "AMD64 and ARM64 Image Build"
                sh "cd $project_path && make build-arm64-by-jenkins IMAGE_TAG=${modbus_lib_version}"
            }
	}
			
        withCredentials([usernamePassword(credentialsId: 'artifactory', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                sh "docker login $artifactory_url --username $USERNAME --password $PASSWORD"
                stage('Artifactory: ARM64') {
                     arm64_local_image="arm64/${image_name}:${modbus_lib_version}"
                     failure_step = "Artifactory: ARM64"
		     sh "docker tag ${arm64_local_image} $artifactory_url/repository/pandora/arm64/${image_name}:${modbus_lib_version}-${image_tag}"
                     sh "docker push $artifactory_url/repository/pandora/arm64/${image_name}:${modbus_lib_version}-${image_tag}"
		     sh "docker tag ${arm64_local_image} $artifactory_url/repository/pandora/arm64/${image_name}:latest"
		     sh "docker push $artifactory_url/repository/pandora/arm64/${image_name}:latest"
                }
                stage('Artifactory: AMD64') {
                     amd64_local_image="amd64/${image_name}:${modbus_lib_version}"
                     failure_step = "Artifactory: AMD64"
		     sh "docker tag ${amd64_local_image} $artifactory_url/repository/pandora/amd64/${image_name}:${modbus_lib_version}-${image_tag}"
                     sh "docker push $artifactory_url/repository/pandora/amd64/${image_name}:${modbus_lib_version}-${image_tag}"
		     sh "docker tag ${amd64_local_image} $artifactory_url/repository/pandora/amd64/${image_name}:latest"
		     sh "docker push $artifactory_url/repository/pandora/amd64/${image_name}:latest"
                }
	}

        currentBuild.result = "SUCCESS"
    }

    catch (Exception e) {
        echo e.message
        currentBuild.result = "FAILURE"
    }

    finally {
        sleep(30)

	    env.sonar_qg = ':white_check_mark: PASSED'

        if (currentBuild.result == "FAILURE") {
            slackSend channel: 'pandora_build', message: "${project_name} - Build failed on step ${failure_step}. Changes made by @${commit_author}. Code changes were ${commit_message}. Results at ${env.BUILD_URL} :hankey:", teamDomain: 'hitachivantara-eng', tokenCredentialId: 'JenkinsSlackIntegration'
        }
        else if (currentBuild.result == "UNSTABLE") {
            slackSend channel: 'pandora_build', message: "${project_name} - Build unstable. Changes made by @${commit_author}. Code changes were ${commit_message}. Results at ${env.BUILD_URL} :face_with_monocle:", teamDomain: 'hitachivantara-eng', tokenCredentialId: 'JenkinsSlackIntegration'
        }
        else if (currentBuild.result == "SUCCESS") {
            slackSend channel: 'pandora_build', message: """
*HIOTA-CASSANDRA* Service - Build succeeded. :donegreen:
:jenkins: *Jenkins* Console Build Results: <${env.BUILD_URL}|*Click here!*>
:sonar: *SonarQube Quality Gate:* ${env.sonar_qg}
:sonar: *SonarQube* Scan Analysis: <https://sonar.orl.eng.hitachivantara.com/dashboard?id=Lumada-Edge-Hiota-Cassandra|*Click here!*>
:blackduck: *BlackDuck* Scan: <https://earth.hitachivantara.com/api/projects/8bd4f65f-e860-4887-9c14-3a1309798c8a/versions/6990ea40-0c89-477e-b1c2-6bf62018ad59/components|*Click here!*>
	    """,
		teamDomain: 'hitachivantara-eng', tokenCredentialId: 'JenkinsSlackIntegration'
        }
        cleanWs()
    }
}
