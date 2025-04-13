
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 8b1b133f5ca5b86ccfe5ba4a767f20687ecb6d4b
kustomize build ./kustomize/overlays/dev
```