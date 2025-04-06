
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 33126c2523f158ebd27497715178422eff707d53
kustomize build ./kustomize/overlays/test
```