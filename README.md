
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 75625ee913639c7a5ae6ac0b71c47a494433c6cc
kustomize build ./kustomize/overlays/test
```