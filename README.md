
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 1f1d4c15a286339eb1103f60113f44e1aac5f55e
kustomize build ./kustomize/overlays/dev
```