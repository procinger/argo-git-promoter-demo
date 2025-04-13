
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 49f11ab8be82f90dfaf9bc196ddb0a6cf1531ea6
kustomize build ./kustomize/overlays/dev
```