
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 69c1685eb41ef52aa2252698a03ed10f0ee8be42
kustomize build ./kustomize/overlays/dev
```