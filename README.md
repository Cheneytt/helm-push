# helm-push

`helm-push` is a [helm](https://github.com/kubernetes/helm) plugin that allows you to push chart package to Nexus

Your nexus should be installed [nexus-repository-helm](https://github.com/sonatype-nexus-community/nexus-repository-helm) plugin.

## Installation

Install the latest version:
```shell
$ helm plugin install https://github.com/Cheneytt/helm-push.git
```

Install a specific version:
```shell
$ helm plugin install https://github.com/Cheneytt/helm-push --version 0.1.0
```

## Quick start

```shell
# Add your chart repository
$ helm repo add repo-name https://yournexushost/repository/yourhelmhost/ --username yourusername --password yourpassword

# Push a chart package to your repository
$ helm push chartname.tar.gz reponame

# Push a chart directory to your repository
$ helm push . reponame

# Update Helm cache
$ helm repo update

# Fetch the chart
$ helm fetch reponame/chartname
```