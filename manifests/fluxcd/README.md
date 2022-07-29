# fluxcd

This addon is built based [FluxCD](https://fluxcd.io/)

## install

```shell
$ helm repo add nginx http://172.24.4.178:9999/helm/ 
$ helm fetch nginx/fluxcd --version 0.0.1

$  helm install fluxcd nginx/fluxcd -n flux-system
```

## X-Definitions

Enable fluxcd addon to use these X-definitions

- [helm](https://kubevela.io/docs/end-user/components/helm) helps to deploy a helm chart from everywhere:
git repo / helm repo / S3 compatible bucket.

- [kustomize](https://kubevela.io/docs/end-user/components/kustomize) helps to deploy a kustomize style artifact.
