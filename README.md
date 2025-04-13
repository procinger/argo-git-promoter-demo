
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0b85ca28875e689fa7bf8ddd4b0d5ac147eb8c34
kustomize build ./kustomize/overlays/test
```