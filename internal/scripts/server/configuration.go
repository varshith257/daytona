// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package server

var Configuration = `{"configuration":{"name":"Daytona Workspace Project","build":{"dockerfile":"Dockerfile"},"features":{"ghcr.io/devcontainers/features/common-utils:1":{"installZsh":"true","username":"daytona","uid":"1000","gid":"1000","upgradePackages":"false"},"ghcr.io/devcontainers/features/docker-in-docker:2":{"version":"24.0.2","moby":false,"dockerDashComposeVersion":"v2"},"ghcr.io/devcontainers/features/go:1":{},"ghcr.io/devcontainers/features/node:1":{},"ghcr.io/devcontainers-contrib/features/typescript:2":{},"ghcr.io/wxw-matt/devcontainer-features/command_runner:latest":{"command1":"npm install -g @devcontainers/cli"}},"overrideFeatureInstallOrder":["ghcr.io/devcontainers/features/common-utils:1","ghcr.io/devcontainers/features/docker-in-docker:2","ghcr.io/devcontainers/features/go:1","ghcr.io/devcontainers/features/node:1","ghcr.io/devcontainers-contrib/features/typescript:2","ghcr.io/wxw-matt/devcontainer-features/command_runner:latest"],"configFilePath":{"$mid":1,"fsPath":"/home/vedran/Projects/supervisor/hack/project_image/.devcontainer/devcontainer.json","path":"/home/vedran/Projects/supervisor/hack/project_image/.devcontainer/devcontainer.json","scheme":"vscode-fileHost"}},"workspace":{"workspaceFolder":"/workspaces/supervisor/hack/project_image","workspaceMount":"type=bind,source=/home/vedran/Projects/supervisor,target=/workspaces/supervisor"},"featuresConfiguration":{"featureSets":[{"sourceInformation":{"type":"oci","manifest":{"schemaVersion":2,"mediaType":"application/vnd.oci.image.manifest.v1+json","config":{"mediaType":"application/vnd.devcontainers","digest":"sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","size":0},"layers":[{"mediaType":"application/vnd.devcontainers.layer.v1+tar","digest":"sha256:325e7207351a55c59c4cdcbab9d39d400b7132cb31567ea5cf9c12cfe7eba5b3","size":27648,"annotations":{"org.opencontainers.image.title":"devcontainer-feature-common-utils.tgz"}}],"annotations":{"com.github.package.type":"devcontainer_feature"}},"manifestDigest":"sha256:0b0b3e64a9ed55be10902c8845778727978ab234d3b3a19b42aac54101a9cec7","featureRef":{"id":"common-utils","owner":"devcontainers","namespace":"devcontainers/features","registry":"ghcr.io","resource":"ghcr.io/devcontainers/features/common-utils","path":"devcontainers/features/common-utils","version":"1","tag":"1"},"userFeatureId":"ghcr.io/devcontainers/features/common-utils:1","userFeatureIdWithoutVersion":"ghcr.io/devcontainers/features/common-utils"},"features":[{"id":"common-utils","version":"1.2.0","name":"Common Debian Utilities","documentationURL":"https://github.com/devcontainers/features/tree/main/src/common-utils","description":"Installs a set of common command line utilities, Oh My Zsh!, and sets up a non-root user.","options":{"installZsh":{"type":"boolean","default":true,"description":"Install ZSH?"},"configureZshAsDefaultShell":{"type":"boolean","default":false,"description":"Change default shell to ZSH?"},"installOhMyZsh":{"type":"boolean","default":true,"description":"Install Oh My Zsh!?"},"upgradePackages":{"type":"boolean","default":true,"description":"Upgrade OS packages?"},"username":{"type":"string","proposals":["vscode","codespace","none","automatic"],"default":"automatic","description":"Enter name of non-root user to configure or none to skip"},"uid":{"type":"string","proposals":["1000","automatic"],"default":"automatic","description":"Enter uid for non-root user"},"gid":{"type":"string","proposals":["1000","automatic"],"default":"automatic","description":"Enter gid for non-root user"},"nonFreePackages":{"type":"boolean","default":false,"description":"Add packages from non-free Debian repository?"}},"included":true,"value":{"installZsh":"true","username":"daytona","uid":"1000","gid":"1000","upgradePackages":"false"},"cachePath":"/tmp/devcontainercli-vedran/container-features/0.55.0-1703800437258/common-utils_0","consecutiveId":"common-utils_0"}],"internalVersion":"2","computedDigest":"sha256:0b0b3e64a9ed55be10902c8845778727978ab234d3b3a19b42aac54101a9cec7"},{"sourceInformation":{"type":"oci","manifest":{"schemaVersion":2,"mediaType":"application/vnd.oci.image.manifest.v1+json","config":{"mediaType":"application/vnd.devcontainers","digest":"sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","size":0},"layers":[{"mediaType":"application/vnd.devcontainers.layer.v1+tar","digest":"sha256:c5335030bde464fb5dfb7ca7aaa5def4f968ac36a786f11fbf3eda4f2dd1eaf0","size":30720,"annotations":{"org.opencontainers.image.title":"devcontainer-feature-docker-in-docker.tgz"}}],"annotations":{"dev.containers.metadata":"{\"id\":\"docker-in-docker\",\"version\":\"2.7.1\",\"name\":\"Docker (Docker-in-Docker)\",\"documentationURL\":\"https://github.com/devcontainers/features/tree/main/src/docker-in-docker\",\"description\":\"Create child containers *inside* a container, independent from the host's docker instance. Installs Docker extension in the container along with needed CLIs.\",\"options\":{\"version\":{\"type\":\"string\",\"proposals\":[\"latest\",\"none\",\"20.10\"],\"default\":\"latest\",\"description\":\"Select or enter a Docker/Moby Engine version. (Availability can vary by OS version.)\"},\"moby\":{\"type\":\"boolean\",\"default\":true,\"description\":\"Install OSS Moby build instead of Docker CE\"},\"dockerDashComposeVersion\":{\"type\":\"string\",\"enum\":[\"none\",\"v1\",\"v2\"],\"default\":\"v1\",\"description\":\"Default version of Docker Compose (v1 or v2 or none)\"},\"azureDnsAutoDetection\":{\"type\":\"boolean\",\"default\":true,\"description\":\"Allow automatically setting the dockerd DNS server when the installation script detects it is running in Azure\"},\"dockerDefaultAddressPool\":{\"type\":\"string\",\"default\":\"\",\"proposals\":[],\"description\":\"Define default address pools for Docker networks. e.g. base=192.168.0.0/16,size=24\"},\"installDockerBuildx\":{\"type\":\"boolean\",\"default\":true,\"description\":\"Install Docker Buildx\"}},\"entrypoint\":\"/usr/local/share/docker-init.sh\",\"privileged\":true,\"containerEnv\":{\"DOCKER_BUILDKIT\":\"1\"},\"customizations\":{\"vscode\":{\"extensions\":[\"ms-azuretools.vscode-docker\"]}},\"mounts\":[{\"source\":\"dind-var-lib-docker-${devcontainerId}\",\"target\":\"/var/lib/docker\",\"type\":\"volume\"}],\"installsAfter\":[\"ghcr.io/devcontainers/features/common-utils\"]}","com.github.package.type":"devcontainer_feature"}},"manifestDigest":"sha256:f6a73ee06601d703db7d95d03e415cab229e78df92bb5002e8559bcfc047fec6","featureRef":{"id":"docker-in-docker","owner":"devcontainers","namespace":"devcontainers/features","registry":"ghcr.io","resource":"ghcr.io/devcontainers/features/docker-in-docker","path":"devcontainers/features/docker-in-docker","version":"2","tag":"2"},"userFeatureId":"ghcr.io/devcontainers/features/docker-in-docker:2","userFeatureIdWithoutVersion":"ghcr.io/devcontainers/features/docker-in-docker"},"features":[{"id":"docker-in-docker","version":"2.7.1","name":"Docker (Docker-in-Docker)","documentationURL":"https://github.com/devcontainers/features/tree/main/src/docker-in-docker","description":"Create child containers *inside* a container, independent from the host's docker instance. Installs Docker extension in the container along with needed CLIs.","options":{"version":{"type":"string","proposals":["latest","none","20.10"],"default":"latest","description":"Select or enter a Docker/Moby Engine version. (Availability can vary by OS version.)"},"moby":{"type":"boolean","default":true,"description":"Install OSS Moby build instead of Docker CE"},"dockerDashComposeVersion":{"type":"string","enum":["none","v1","v2"],"default":"v1","description":"Default version of Docker Compose (v1 or v2 or none)"},"azureDnsAutoDetection":{"type":"boolean","default":true,"description":"Allow automatically setting the dockerd DNS server when the installation script detects it is running in Azure"},"dockerDefaultAddressPool":{"type":"string","default":"","proposals":[],"description":"Define default address pools for Docker networks. e.g. base=192.168.0.0/16,size=24"},"installDockerBuildx":{"type":"boolean","default":true,"description":"Install Docker Buildx"}},"entrypoint":"/usr/local/share/docker-init.sh","privileged":true,"containerEnv":{"DOCKER_BUILDKIT":"1"},"customizations":{"vscode":{"extensions":["ms-azuretools.vscode-docker"]}},"mounts":[{"source":"dind-var-lib-docker-${devcontainerId}","target":"/var/lib/docker","type":"volume"}],"installsAfter":["ghcr.io/devcontainers/features/common-utils"],"included":true,"value":{"version":"24.0.2","moby":false,"dockerDashComposeVersion":"v2"},"cachePath":"/tmp/devcontainercli-vedran/container-features/0.55.0-1703800437258/docker-in-docker_1","consecutiveId":"docker-in-docker_1"}],"internalVersion":"2","computedDigest":"sha256:f6a73ee06601d703db7d95d03e415cab229e78df92bb5002e8559bcfc047fec6"},{"sourceInformation":{"type":"oci","manifest":{"schemaVersion":2,"mediaType":"application/vnd.oci.image.manifest.v1+json","config":{"mediaType":"application/vnd.devcontainers","digest":"sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","size":0},"layers":[{"mediaType":"application/vnd.devcontainers.layer.v1+tar","digest":"sha256:6ad459521274db1a3b374a2cc5fc693b099d9c2caa433474dc69c9148d69a015","size":16384,"annotations":{"org.opencontainers.image.title":"devcontainer-feature-go.tgz"}}],"annotations":{"dev.containers.metadata":"{\"id\":\"go\",\"version\":\"1.2.2\",\"name\":\"Go\",\"documentationURL\":\"https://github.com/devcontainers/features/tree/main/src/go\",\"description\":\"Installs Go and common Go utilities. Auto-detects latest version and installs needed dependencies.\",\"options\":{\"version\":{\"type\":\"string\",\"proposals\":[\"latest\",\"none\",\"1.21\",\"1.20\"],\"default\":\"latest\",\"description\":\"Select or enter a Go version to install\"},\"golangciLintVersion\":{\"type\":\"string\",\"default\":\"latest\",\"description\":\"Version of golangci-lint to install\"}},\"init\":true,\"customizations\":{\"vscode\":{\"extensions\":[\"golang.Go\"]}},\"containerEnv\":{\"GOROOT\":\"/usr/local/go\",\"GOPATH\":\"/go\",\"PATH\":\"/usr/local/go/bin:/go/bin:${PATH}\"},\"capAdd\":[\"SYS_PTRACE\"],\"securityOpt\":[\"seccomp=unconfined\"],\"installsAfter\":[\"ghcr.io/devcontainers/features/common-utils\"]}","com.github.package.type":"devcontainer_feature"}},"manifestDigest":"sha256:d5e34970f31942a4d9f6b3f6f52da1176b2dcb35aeba3f0fc93b974e287d3b16","featureRef":{"id":"go","owner":"devcontainers","namespace":"devcontainers/features","registry":"ghcr.io","resource":"ghcr.io/devcontainers/features/go","path":"devcontainers/features/go","version":"1","tag":"1"},"userFeatureId":"ghcr.io/devcontainers/features/go:1","userFeatureIdWithoutVersion":"ghcr.io/devcontainers/features/go"},"features":[{"id":"go","version":"1.2.2","name":"Go","documentationURL":"https://github.com/devcontainers/features/tree/main/src/go","description":"Installs Go and common Go utilities. Auto-detects latest version and installs needed dependencies.","options":{"version":{"type":"string","proposals":["latest","none","1.21","1.20"],"default":"latest","description":"Select or enter a Go version to install"},"golangciLintVersion":{"type":"string","default":"latest","description":"Version of golangci-lint to install"}},"init":true,"customizations":{"vscode":{"extensions":["golang.Go"]}},"containerEnv":{"GOROOT":"/usr/local/go","GOPATH":"/go","PATH":"/usr/local/go/bin:/go/bin:${PATH}"},"capAdd":["SYS_PTRACE"],"securityOpt":["seccomp=unconfined"],"installsAfter":["ghcr.io/devcontainers/features/common-utils"],"included":true,"value":{},"cachePath":"/tmp/devcontainercli-vedran/container-features/0.55.0-1703800437258/go_2","consecutiveId":"go_2"}],"internalVersion":"2","computedDigest":"sha256:d5e34970f31942a4d9f6b3f6f52da1176b2dcb35aeba3f0fc93b974e287d3b16"},{"sourceInformation":{"type":"oci","manifest":{"schemaVersion":2,"mediaType":"application/vnd.oci.image.manifest.v1+json","config":{"mediaType":"application/vnd.devcontainers","digest":"sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","size":0},"layers":[{"mediaType":"application/vnd.devcontainers.layer.v1+tar","digest":"sha256:d849e1ca911a1acbb591d5b29b06ee98b70644aa3845a8b5ce83a96677030c68","size":18944,"annotations":{"org.opencontainers.image.title":"devcontainer-feature-node.tgz"}}],"annotations":{"dev.containers.metadata":"{\"id\":\"node\",\"version\":\"1.3.1\",\"name\":\"Node.js (via nvm), yarn and pnpm\",\"documentationURL\":\"https://github.com/devcontainers/features/tree/main/src/node\",\"description\":\"Installs Node.js, nvm, yarn, pnpm, and needed dependencies.\",\"options\":{\"version\":{\"type\":\"string\",\"proposals\":[\"lts\",\"latest\",\"none\",\"18\",\"16\",\"14\"],\"default\":\"lts\",\"description\":\"Select or enter a Node.js version to install\"},\"nodeGypDependencies\":{\"type\":\"boolean\",\"default\":true,\"description\":\"Install dependencies to compile native node modules (node-gyp)?\"},\"nvmInstallPath\":{\"type\":\"string\",\"default\":\"/usr/local/share/nvm\",\"description\":\"The path where NVM will be installed.\"},\"nvmVersion\":{\"type\":\"string\",\"proposals\":[\"latest\",\"0.39\"],\"default\":\"latest\",\"description\":\"Version of NVM to install.\"}},\"customizations\":{\"vscode\":{\"extensions\":[\"dbaeumer.vscode-eslint\"]}},\"containerEnv\":{\"NVM_DIR\":\"/usr/local/share/nvm\",\"NVM_SYMLINK_CURRENT\":\"true\",\"PATH\":\"/usr/local/share/nvm/current/bin:${PATH}\"},\"installsAfter\":[\"ghcr.io/devcontainers/features/common-utils\"]}","com.github.package.type":"devcontainer_feature"}},"manifestDigest":"sha256:7d31b83459dd5110c37e7f5acb2920335cb1e5ebf014326d7eb6a0b290cc820a","featureRef":{"id":"node","owner":"devcontainers","namespace":"devcontainers/features","registry":"ghcr.io","resource":"ghcr.io/devcontainers/features/node","path":"devcontainers/features/node","version":"1","tag":"1"},"userFeatureId":"ghcr.io/devcontainers/features/node:1","userFeatureIdWithoutVersion":"ghcr.io/devcontainers/features/node"},"features":[{"id":"node","version":"1.3.1","name":"Node.js (via nvm), yarn and pnpm","documentationURL":"https://github.com/devcontainers/features/tree/main/src/node","description":"Installs Node.js, nvm, yarn, pnpm, and needed dependencies.","options":{"version":{"type":"string","proposals":["lts","latest","none","18","16","14"],"default":"lts","description":"Select or enter a Node.js version to install"},"nodeGypDependencies":{"type":"boolean","default":true,"description":"Install dependencies to compile native node modules (node-gyp)?"},"nvmInstallPath":{"type":"string","default":"/usr/local/share/nvm","description":"The path where NVM will be installed."},"nvmVersion":{"type":"string","proposals":["latest","0.39"],"default":"latest","description":"Version of NVM to install."}},"customizations":{"vscode":{"extensions":["dbaeumer.vscode-eslint"]}},"containerEnv":{"NVM_DIR":"/usr/local/share/nvm","NVM_SYMLINK_CURRENT":"true","PATH":"/usr/local/share/nvm/current/bin:${PATH}"},"installsAfter":["ghcr.io/devcontainers/features/common-utils"],"included":true,"value":{},"cachePath":"/tmp/devcontainercli-vedran/container-features/0.55.0-1703800437258/node_3","consecutiveId":"node_3"}],"internalVersion":"2","computedDigest":"sha256:7d31b83459dd5110c37e7f5acb2920335cb1e5ebf014326d7eb6a0b290cc820a"},{"sourceInformation":{"type":"oci","manifest":{"schemaVersion":2,"mediaType":"application/vnd.oci.image.manifest.v1+json","config":{"mediaType":"application/vnd.devcontainers","digest":"sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","size":0},"layers":[{"mediaType":"application/vnd.devcontainers.layer.v1+tar","digest":"sha256:1aaa167d2b7921b33d1ef90e29f390a79875613c4095f0893c8450a609714754","size":12288,"annotations":{"org.opencontainers.image.title":"devcontainer-feature-typescript.tgz"}}],"annotations":{"com.github.package.type":"devcontainer_feature"}},"manifestDigest":"sha256:2e7e3e0b17a0c7dbdf78bf0cb9ee2cc4a0a4d53c6ad96cadadae167bc2885ed8","featureRef":{"id":"typescript","owner":"devcontainers-contrib","namespace":"devcontainers-contrib/features","registry":"ghcr.io","resource":"ghcr.io/devcontainers-contrib/features/typescript","path":"devcontainers-contrib/features/typescript","version":"2","tag":"2"},"userFeatureId":"ghcr.io/devcontainers-contrib/features/typescript:2","userFeatureIdWithoutVersion":"ghcr.io/devcontainers-contrib/features/typescript"},"features":[{"id":"typescript","version":"2.0.14","name":"TypeScript (via npm)","documentationURL":"http://github.com/devcontainers-contrib/features/tree/main/src/typescript","description":"TypeScript is a strongly typed programming language that builds on JavaScript, giving you better tooling at any scale.","options":{"version":{"default":"latest","description":"Select the version to install.","proposals":["latest"],"type":"string"}},"installsAfter":["ghcr.io/devcontainers-contrib/features/npm-package"],"included":true,"value":{},"cachePath":"/tmp/devcontainercli-vedran/container-features/0.55.0-1703800437258/typescript_4","consecutiveId":"typescript_4"}],"internalVersion":"2","computedDigest":"sha256:2e7e3e0b17a0c7dbdf78bf0cb9ee2cc4a0a4d53c6ad96cadadae167bc2885ed8"},{"sourceInformation":{"type":"oci","manifest":{"schemaVersion":2,"mediaType":"application/vnd.oci.image.manifest.v1+json","config":{"mediaType":"application/vnd.devcontainers","digest":"sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","size":0},"layers":[{"mediaType":"application/vnd.devcontainers.layer.v1+tar","digest":"sha256:80b7c3fa6f7aaa2dc17cfeb2cfa268e5abad1daecfb3d10e888c1b313bbaa43c","size":7168,"annotations":{"org.opencontainers.image.title":"devcontainer-feature-command_runner.tgz"}}],"annotations":{"com.github.package.type":"devcontainer_feature"}},"manifestDigest":"sha256:ee54b2f7a82b373be5ca093726187a4c1bfa604c0ac3e07e787e68f31fca2736","featureRef":{"id":"command_runner","owner":"wxw-matt","namespace":"wxw-matt/devcontainer-features","registry":"ghcr.io","resource":"ghcr.io/wxw-matt/devcontainer-features/command_runner","path":"wxw-matt/devcontainer-features/command_runner","version":"latest","tag":"latest"},"userFeatureId":"ghcr.io/wxw-matt/devcontainer-features/command_runner:latest","userFeatureIdWithoutVersion":"ghcr.io/wxw-matt/devcontainer-features/command_runner"},"features":[{"name":"Run shell commands","id":"command_runner","version":"0.5.0","description":"A feature to run any valid shell commands, such as apt-get, curl/wget, of course, bash/zsh","options":{"command1":{"type":"string","default":"","description":"The command 1."},"command2":{"type":"string","default":"","description":"The command 2."},"command3":{"type":"string","default":"","description":"The command 3."}},"installsAfter":[],"included":true,"value":{"command1":"npm install -g @devcontainers/cli"},"cachePath":"/tmp/devcontainercli-vedran/container-features/0.55.0-1703800437258/command_runner_5","consecutiveId":"command_runner_5"}],"internalVersion":"2","computedDigest":"sha256:ee54b2f7a82b373be5ca093726187a4c1bfa604c0ac3e07e787e68f31fca2736"}],"dstFolder":"/tmp/devcontainercli-vedran/container-features/0.55.0-1703800437258"},"mergedConfiguration":{"name":"Daytona Workspace Project","build":{"dockerfile":"Dockerfile"},"features":{"ghcr.io/devcontainers/features/common-utils:1":{"installZsh":"true","username":"daytona","uid":"1000","gid":"1000","upgradePackages":"false"},"ghcr.io/devcontainers/features/docker-in-docker:2":{"version":"24.0.2","moby":false,"dockerDashComposeVersion":"v2"},"ghcr.io/devcontainers/features/go:1":{},"ghcr.io/devcontainers/features/node:1":{},"ghcr.io/devcontainers-contrib/features/typescript:2":{},"ghcr.io/wxw-matt/devcontainer-features/command_runner:latest":{"command1":"npm install -g @devcontainers/cli"}},"overrideFeatureInstallOrder":["ghcr.io/devcontainers/features/common-utils:1","ghcr.io/devcontainers/features/docker-in-docker:2","ghcr.io/devcontainers/features/go:1","ghcr.io/devcontainers/features/node:1","ghcr.io/devcontainers-contrib/features/typescript:2","ghcr.io/wxw-matt/devcontainer-features/command_runner:latest"],"configFilePath":{"$mid":1,"fsPath":"/home/vedran/Projects/supervisor/hack/project_image/.devcontainer/devcontainer.json","path":"/home/vedran/Projects/supervisor/hack/project_image/.devcontainer/devcontainer.json","scheme":"vscode-fileHost"},"init":true,"privileged":true,"capAdd":["SYS_PTRACE"],"securityOpt":["seccomp=unconfined"],"entrypoints":["/usr/local/share/docker-init.sh"],"mounts":[{"source":"dind-var-lib-docker-${devcontainerId}","target":"/var/lib/docker","type":"volume"}],"customizations":{"vscode":[{"extensions":["ms-azuretools.vscode-docker"]},{"extensions":["golang.Go"]},{"extensions":["dbaeumer.vscode-eslint"]}]},"onCreateCommands":[],"updateContentCommands":[],"postCreateCommands":[],"postStartCommands":[],"postAttachCommands":[],"remoteEnv":{},"containerEnv":{},"portsAttributes":{}}}`