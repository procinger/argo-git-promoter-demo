
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 2c0018e9fa0c3afaea900e16c6e191e73286c5da
kustomize build ./kustomize/overlays/dev
```