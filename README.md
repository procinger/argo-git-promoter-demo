
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0d2863fa89180df484ff9ea273ced8c6e7d5f6a0
kustomize build ./kustomize/overlays/dev
```