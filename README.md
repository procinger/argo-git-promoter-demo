
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout d49dfdad505d62dbeae47ea5a246eca986e2203c
kustomize build ./kustomize/overlays/test
```